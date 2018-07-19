/*
   Copyright 2018 The HAWQ Team.
*/

package myresource

import (
	"log"

	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"

	"fmt"
	"time"

	"github.com/golang/glog"
	"github.com/hawq-cn/apiserver-example/pkg/apis/core/v1alpha1"
	listers "github.com/hawq-cn/apiserver-example/pkg/client/listers_generated/core/v1alpha1"
	"github.com/hawq-cn/apiserver-example/pkg/controller/sharedinformers"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

// +controller:group=core,version=v1alpha1,kind=MyResource,resource=myresources
type MyResourceControllerImpl struct {
	builders.DefaultControllerFns

	// lister indexes properties about MyResource
	lister listers.MyResourceLister

	// informer
	informerSynced cache.InformerSynced
	queue          workqueue.RateLimitingInterface
	threadiness    int
}

// Init initializes the controller and is called by the generated code
// Register watches for additional resource types here.
func (c *MyResourceControllerImpl) Init(arguments sharedinformers.ControllerInitArguments) {
	log.Printf("my controller is initing\n")

	// Use the lister for indexing myresources labels
	c.lister = arguments.GetSharedInformers().Factory.Core().V1alpha1().MyResources().Lister()

	my_informer := arguments.GetSharedInformers().Factory.Core().V1alpha1().MyResources().Informer()

	my_informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    c.addFunc,
		UpdateFunc: c.updateFunc,
		DeleteFunc: c.deleteFunc,
	})

	c.informerSynced = my_informer.HasSynced
	c.queue = workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "my_resource_queue")
	c.threadiness = 4
}

// Reconcile handles enqueued messages
func (c *MyResourceControllerImpl) Reconcile(u *v1alpha1.MyResource) error {
	// Implement controller logic here
	log.Printf("Running reconcile MyResource for %s\n", u.Name)
	return nil
}

func (c *MyResourceControllerImpl) Get(namespace, name string) (*v1alpha1.MyResource, error) {
	return c.lister.MyResources(namespace).Get(name)
}

func (c *MyResourceControllerImpl) Run(stopCh <-chan struct{}) {
	log.Printf("my controller is running\n")

	if !cache.WaitForCacheSync(stopCh, c.informerSynced) {
		runtime.HandleError(fmt.Errorf("timed out waiting for caches to sync"))
		return
	}
	log.Printf("informer cache synced\n")

	for i := 0; i < c.threadiness; i++ {
		go wait.Until(c.runWorker, time.Second, stopCh)
	}

	<-stopCh
}

func (c *MyResourceControllerImpl) sync(key string) error {
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		runtime.HandleError(fmt.Errorf("invalid resource key: %s", key))
		return nil
	}

	// current event
	myRes, err := c.lister.MyResources(namespace).Get(name)
	if err != nil {
		if apierrors.IsNotFound(err) {
			// hit delete event
		} else {
			return err
		}
	} else {
		if myRes != nil {
			log.Printf("in sync handler, resource: %v", myRes)
		}
	}

	// list all
	selector, _ := labels.Parse("")
	arr, _ := c.lister.MyResources(namespace).List(selector)
	var resNames string = ""
	for i := 0; i < len(arr); i++ {
		resNames += arr[i].Name + " "
	}
	log.Printf("There are %d resource(s): %s\n", len(arr), resNames)

	return nil
}

func (c *MyResourceControllerImpl) runWorker() {
	for c.processNextWorkItem() {
	}
}

// processNextWorkItem deals with one key off the queue.  It returns false when it's time to quit.
func (c *MyResourceControllerImpl) processNextWorkItem() bool {
	key, quit := c.queue.Get()
	if quit {
		return false
	}
	defer c.queue.Done(key)

	err := c.sync(key.(string))
	if err == nil {
		c.queue.Forget(key)
		return true
	}

	runtime.HandleError(fmt.Errorf("%v failed with: %v", key, err))
	c.queue.AddRateLimited(key)

	return true
}

func (c *MyResourceControllerImpl) enqueue(obj *v1alpha1.MyResource) {
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
	if err != nil {
		runtime.HandleError(fmt.Errorf("Couldn't get key for object %#v: %v", obj, err))
		return
	}

	c.queue.Add(key)
}

// informer add/update/delete handler

func (c *MyResourceControllerImpl) addFunc(obj interface{}) {
	castObj := obj.(*v1alpha1.MyResource)
	log.Printf("Adding resource %s", castObj.Name)
	c.enqueue(castObj)
}

func (c *MyResourceControllerImpl) updateFunc(obj, _ interface{}) {
	castObj := obj.(*v1alpha1.MyResource)
	log.Printf("Updating resource %s", castObj.Name)
	c.enqueue(castObj)
}

func (c *MyResourceControllerImpl) deleteFunc(obj interface{}) {
	castObj, ok := obj.(*v1alpha1.MyResource)
	if !ok {
		tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
		if !ok {
			glog.Errorf("Couldn't get object from tombstone %#v", obj)
			return
		}
		castObj, ok = tombstone.Obj.(*v1alpha1.MyResource)
		if !ok {
			glog.Errorf("Tombstone contained object that is not expected %#v", obj)
			return
		}
	}
	log.Printf("Deleting resource %s", castObj.Name)
	c.enqueue(castObj)
}
