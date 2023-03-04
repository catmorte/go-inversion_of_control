# go-inversion_of_control

Lib to register and retrieve beans by type from context. Similar to service locator

### How to use:

First of all you should init context by doing the following in your project:

    package appPackage
    import 	"github.com/catmorte/go-inversion_of_control/pkg/context"
    
    func init() {
      context.SetContext(context.NewMemoryContext())
    }

It creates an instance of default context implementation from package:
 
    github.com/catmorte/go-inversion_of_control/pkg/context
    
The next step is to create beans by doing the following in your project:

    package appPackage/beans
    import (
      ...
      "github.com/catmorte/go-inversion_of_control/pkg/context"
    )
    
    func init() {
      // start dependencies definition
      dep1 := context.Dep[*dep1Type]()
      dep2 := context.Dep[*dep1Type]()
      ...
      depN := context.Dep((*depNType)(nil))
      // end dependencies definition


      // start bean constructor definition 
      context.Reg[*beanType](context.GetContext(), func() interface{} {
        return NewBeanType(
                    (<-dep1.Waiter).(*dep1Type),
                    (<-dep2.Waiter).(*dep2Type),
                    ...
                    (<-depN.Waiter).(*depNType),

        			)
      }, dep1, dep2, ..., depN)
      // end bean constructor definition 
    }
    
**Reg** function starts bean initialization and waits until all the necessary dependencies will be initialized.

**! Please ensure that you pass all the deps in the function Reg as well as you using them inside constructor, otherwise it will block bean initialization**

The next step is to import context and beans initialization:

    import (
        ...
         _ "appPackage"        // initialize the context
         _ "appPackage/beans"  // initialize beans
        ...
     )

**! Make sure that the context initialization imported before beans import**

To retrieve bean use the following in your project:

   	bean := context.Ask[*beanType](context.GetContext())

**Ask** function waits until bean will be initialized and then retrieve it

You may find the working example in folder /example

Also supported **named scopes**: **RegScoped, AskScoped, DepScoped**
    
