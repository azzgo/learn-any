import globalState from "./globalState";

export class Dep {
  constructor() {
    this.subs = [];
  }

  addSub(sub) {
    this.subs.push(sub);
  }

  removeSub(sub) {
    remove(this.subs, sub);
  }

  depend() {
    if (globalState.target) {
      this.addSub(globalState.target);
    }
  }

  notify() {
    const subs = this.subs.slice();
    for (let sub of subs) {
      sub.update();
    }
  }

  remove(arr, item) {
    if (arr.length) {
      const index = arr.indexOf(item);
      return arr.splice(index, 1);
    }
  }
}
