import { Observer } from '../src/observer';
import Watcher from '../src/wacher';
import * as assert from 'assert';
let testObj;

beforeEach(() => {
  testObj = {
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
});

test('测试 对象 属性更新', () => {
  new Watcher(testObj, 'b.c', function(newVal) {
    expect(newVal).toBe(3);
  });

  testObj.b.c++;
});

test('测试 数组 更新，所有数组项 都会收到更新', () => {
  new Watcher(testObj, 'g', function(newVal) {
    assert.strictEqual(newVal.length, 5);
  });

  new Watcher(testObj, 'g.1', function(newVal) {
    expect(newVal).toBeDefined();
    expect(newVal).toBe(2);
  });

  testObj.g.push('item');
});

test('新增 对象 的属性更新', () => {
  testObj.g.push({ h: 5 });

  new Watcher(testObj, 'g.4.h', function(newVal) {
    expect(newVal).toBe(6);
  });

  testObj.g[testObj.g.length - 1].h++;
});

