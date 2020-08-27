import globalState from "./globalState";

// 简单的订阅发布模式
export class Dep {
  constructor() {
    this.subs = [];
  }

  _addSub(sub) {
    this.subs.push(sub);
  }

  depend() {
    if (globalState.target) {
      this._addSub(globalState.target);
    }
  }

  notify() {
    const subs = this.subs.slice();
    for (let sub of subs) {
      sub.update();
    }
  }

  // useless code
  removeSub(sub) {
    remove(this.subs, sub);
  }

  // useless code
  remove(arr, item) {
    if (arr.length) {
      const index = arr.indexOf(item);
      return arr.splice(index, 1);
    }
  }
}
