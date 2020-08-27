import { Dep } from './dep';
import { arraryMethods } from './arrayMethodModifier';
import { def } from './utils';

export function defineReactive(data, key, val) {
  let ob = observe(val);
  let dep = new Dep();

  Object.defineProperty(data, key, {
    enumerable: true,
    configurable: true,
    get: function() {
      // 处理属性
      dep.depend();
      ob && ob.dep.depend();
      return val;
    },
    set: function(newVal) {
      if (val === newVal) {
        return;
      }
      val = newVal;
      dep.notify();
    },
  });
}

function observe(value) {
  if (typeof value !== 'object') {
    return;
  }

  let ob;

  if (value.hasOwnProperty('__ob__') && value.__ob__ instanceof Observer) {
    ob = value.__ob__;
  } else {
    ob = new Observer(value);
  }
  return ob;
}



export class Observer {
  constructor(value) {
    this.value = value;
    this.dep = new Dep();
    def(value, '__ob__', this);

    if (Array.isArray(value)) {
      value.__proto__ = arraryMethods;
      this.observeArrary(value)
    } else {
      this.walk(value);
    }
  }

  walk(obj) {
    const keys = Object.keys(obj);
    for (let key of keys) {
      defineReactive(obj, key, obj[key]);
    }
  }

  observeArrary(items) {
    for (let val of items) {
      observe(val)
    }
  }
}
