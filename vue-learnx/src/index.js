import { Observer } from './data-binding/observer';
import Watcher from './data-binding/wacher';
import * as assert from 'assert';

let testObj = {
  a: '1',
  b: {
    c: 2,
    d: {
      e: 3,
      f: 'hello',
    },
  },
  g: [1, 2, 3, 4],
};


new Observer(testObj);


new Watcher(testObj, 'b.c', function(newVal) {
  assert.strictEqual(newVal, 3);
});

new Watcher(testObj, 'g', function(newVal) {
  assert.strictEqual(newVal.length, 5);
});

new Watcher(testObj, 'g.4.h', function(newVal) {
  assert.strictEqual(newVal, 5);
});

new Watcher(testObj, 'g.0', function(newVal) {
  assert.strictEqual(newVal, 1);
});



testObj.b.c++;

testObj.g.push({ h: 5});

testObj.g[4]++;
