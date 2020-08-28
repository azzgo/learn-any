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
 * @description 解析路径, 返回一个对应路径的 getter，getter 会尝试获取传入对象的指定路径下的值
 * @param {string} path "a.1.b" 的形式
 * @returns {(obj: object) => any } 返回传入对象的指定路径下的值, 取不到返回 undefined
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
