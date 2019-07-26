import 'package:sqflite/sqflite.dart';
import 'dart:async';
import 'package:path/path.dart';

import 'Sticky.dart';

import 'dart:developer' as developer;

final tableName = 'STICKIES';

class StickyProvider {
  Database db;

  Future open() async {

    developer.log(await getDatabasesPath());
    
    db = await openDatabase(join(await getDatabasesPath(), 'stickies.db'),
        onCreate: (db, version) {
      db.execute(
          '''
          CREATE TABLE $tableName (id INTEGER PRIMARY KEY, title TEXT, content TEXT, modify_time DATETIME);
          INSERT INTO $tableName (title, content, modify_time) VALUES ("欢迎使用便签", "你可以在便签里快速记录灵感，并支持添加图片、项目列表、待办事件。", "2018-12-20 18:25");
          INSERT INTO $tableName (title, content, modify_time) VALUES ("在氢视窗中查看", "你可以将便签显示在氢视窗中，方便你随时查看。", "2018-12-20 18:25");
          INSERT INTO $tableName (title, content, modify_time) VALUES ("通过图片分享", "你可以将你的便签生成为长图片，方便你分享到设计平台。", "2018-12-20 18:25");
          INSERT INTO $tableName (title, content, modify_time) VALUES ("购物清单", "洗衣液 洗发水 牙膏 鸡蛋", "2017-11-13 07:43");
          ''');
    }, version: 1);
  }

  Future close() async {
    db.close();
  }

  Future<Sticky> insertSticky(Sticky sticky) async {
    sticky.modifyTime = DateTime.now();
    sticky.id = await db.insert(tableName, sticky.toMap());

    return sticky;
  }

  Future<int> updateSticky(Sticky sticky) async {
    return db.update(tableName, sticky.toMap(),
        where: 'id=?', whereArgs: [sticky.id]);
  }

  Future<int> deleteSticky(int id) async {
    return db.delete(tableName, where: 'id=?', whereArgs: [id]);
  }

  Future<Iterable<Sticky>> getStickies() async {
    List<Map> stickies = await db.query(tableName,
        columns: ['title', 'content', 'id', 'modify_time'], where: '1=1');

    return stickies.map((stickyMap) => Sticky.fromMap(stickyMap));
  }

  Future<Sticky> getSticky(int id) async {
    List<Map> stickies = await db.query(tableName,
        columns: ['title', 'content', 'id', 'modify_time'],
        where: 'id=?',
        whereArgs: [id]);

    if (stickies.length > 0) {
      return Sticky.fromMap(stickies.first);
    }

    return null;
  }
}
