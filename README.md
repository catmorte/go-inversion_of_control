# go-inversion_of_control

Lib to register and retrieve beans by type from context. Similar to service locator

### How to use:

First of all you should init context by doing the following in your project:

    package appPackage
    import 	"github.com/catmorte/go-inversion_of_control/pkg/context"
    
    func init() {
      context.SetContext(context.NewMemoryContext())
    }

It create an instance of predefined memory implementation of context in package:
 
    github.com/catmorte/go-inversion_of_control/pkg/context
    
The next step is to create beans by doing the following in your project:

    package appPackage/beans
    import (
      ...
      "github.com/catmorte/go-inversion_of_control/pkg/context"
    )
    
    func init() {
      // start dependencies definition
      dep1 := context.Dep((*dep1Type)(nil))
      dep2 := context.Dep((*dep2Type)(nil))
      ...
      depN := context.Dep((*depNType)(nil))
      // end dependencies definition


      // start bean constructor definition 
      context.GetContext().Reg((*beanType)(nil), func() interface{} {
        return NewBeanType(
                    (<-dep1.Waiter).(*dep1Type),
                    (<-dep2.Waiter).(*dep2Type),
                    ...
                    (<-depN.Waiter).(*depNType),

        			)
      }, dep1, dep2, ..., depN)
      // end bean constructor definition 
    }
    
**Reg** function start registration of bean and wait until all the necessary dependencies will be resolved.

**! Please ensure that you pass all the deps in the function Reg as well as you using them inside constructor, otherwise it will block the constructor**

The next step is to import context initializer and beans initialization:

    import (
        ...
         _ "appPackage"        // initialize the context
         _ "appPackage/beans"  // initialize beans
        ...
     )

**! Make sure that the code is imported before beans import**

To retrieve bean use the following in your project:

   	bean := (<-context.GetContext().Ask((*beanType)(nil))).(*beanType)

**Ask** function waiting until bean resolve all the dependencies and initialize and the retrieve it

You may find the working example in folder /example