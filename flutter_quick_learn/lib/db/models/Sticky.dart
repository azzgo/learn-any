import 'package:intl/intl.dart';

final dateFormat = new DateFormat('yyyy-MM-dd HH:mm:ss');


class Sticky {
  int id;
  String title;
  String content;
  bool isTop;
  DateTime modifyTime;

  Sticky({this.title, this.content}): isTop = false;

  Sticky.fromMap(Map<String, dynamic> stickyMap) {
    id = stickyMap['id'];
    title = stickyMap['title'];
    content = stickyMap['content'];
    isTop = stickyMap["isTop"] == 1;

    try {
      modifyTime = DateTime.parse(stickyMap['modify_time']);
    } catch(e) {}
  }

  Map<String, dynamic> toMap() {
    return {
      "id": id,
      "title": title,
      "content": content,
      "isTop": isTop == true ? 1 : 0,
      "modify_time": dateFormat.format(modifyTime)
    };
  }
}
