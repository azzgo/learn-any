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
    modifyTime = DateTime.parse(stickyMap['modify_time']);
  }

  Map<String, dynamic> toMap() {
    return {
      "id": id,
      "title": title,
      "content": content,
      "modify_time": modifyTime
    };
  }
}
