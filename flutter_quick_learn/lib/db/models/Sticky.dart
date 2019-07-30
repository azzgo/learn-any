import 'package:intl/intl.dart';

final dateFormat = new DateFormat('yyyy-MM-dd HH:mm:ss');


class Sticky {
  int id;
  String title;
  String content;
  DateTime modifyTime;

  Sticky({this.title, this.content});

  Sticky.fromMap(Map<String, dynamic> stickyMap) {
    id = stickyMap['id'];
    title = stickyMap['title'];
    content = stickyMap['content'];

    try {
      modifyTime = DateTime.parse(stickyMap['modify_time']);
    } catch(e) {}
  }

  Map<String, dynamic> toMap() {
    return {
      "id": id,
      "title": title,
      "content": content,
      "modify_time": dateFormat.format(modifyTime)
    };
  }
}
