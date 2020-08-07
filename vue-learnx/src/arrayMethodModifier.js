import { def } from './utils';

const arraryProto = Array.prototype;

export const arraryMethods = Object.create(arraryProto);

['push', 'pop', 'shift', 'unshift', 'splice', 'sort', 'reverse'].forEach(
  method => {
    const originMethod = arraryMethods[method];

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
