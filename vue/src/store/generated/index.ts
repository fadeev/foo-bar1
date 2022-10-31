// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import FadeevFoobar1Foobar1 from './fadeev.foobar1.foobar1'


export default { 
  FadeevFoobar1Foobar1: load(FadeevFoobar1Foobar1, 'fadeev.foobar1.foobar1'),
  
}


function load(mod, fullns) {
    return function init(store) {        
        if (store.hasModule([fullns])) {
            throw new Error('Duplicate module name detected: '+ fullns)
        }else{
            store.registerModule([fullns], mod)
            store.subscribe((mutation) => {
                if (mutation.type == 'common/env/INITIALIZE_WS_COMPLETE') {
                    store.dispatch(fullns+ '/init', null, {
                        root: true
                    })
                }
            })
        }
    }
}