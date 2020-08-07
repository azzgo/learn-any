import { Observer } from './observer';
import Watcher from './wacher';
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
  // assert.strictEqual(newVal, 6);
  console.log('g.4.h:' + newVal)
});

new Watcher(testObj, 'g.0', function(newVal) {
  // assert.strictEqual(newVal, 6);
  console.log('g.0:' + newVal)
});


testObj.b.c++;

testObj.g.push({ h: 5});
