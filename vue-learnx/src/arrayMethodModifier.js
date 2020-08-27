import { def } from './utils';

const arrayProto = Array.prototype;

// 拷贝一份 Array 的原型方法
export const arraryMethods = Object.create(arrayProto);

['push', 'pop', 'shift', 'unshift', 'splice', 'sort', 'reverse'].forEach(
  method => {
    const originMethod = arraryMethods[method];

    // 改写拷贝的 Array 原型方法
    def(arraryMethods, method, function mutator(...args) {
      const result = originMethod.apply(this, args);
      const ob = this.__ob__;
      let inserted;
      switch (method) {
        case 'push':
        case 'unshift':
          inserted = args;
          break;
        case 'splice':
          inserted = args.slice(2);
          break;
      }
      if (inserted) ob.observeArrary(inserted);
      ob.dep.notify();
      return result;
    });
  },
);
