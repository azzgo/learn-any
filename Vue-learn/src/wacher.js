import globalState from './globalState'

export default class Watcher {
  constructor(vm, exp, cb) {
    this.vm = vm
    this.getter = parsePath(exp);
    this.cb = cb;
    this.value = this.get();
  }

  get() {
    globalState.target = this;
    let value = this.getter.call(null, this.vm);
    globalState.target = undefined;
    return value;
  }

  update() {
    const oldValue = this.value;
    this.value = this.get();
    this.cb.call(null, this.value, oldValue);
  }
}

const bailRE = /[^\w.$]/;
/**
 * @param {string} path
 */
function parsePath(path) {
  if (bailRE.test(path)) {
    return;
  }

  const segments = path.split('.');
  return function(obj) {
    for (let segment of segments) {
      if (!obj) return;
      obj = obj[segment];
    }
    return obj;
  };
}
