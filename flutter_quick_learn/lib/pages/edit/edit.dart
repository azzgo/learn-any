import 'package:flutter/material.dart';

class EditStickyPage extends StatefulWidget {
  int id;
  String title;
  String content;

  EditStickyPage({this.id, this.title, this.content});

  @override
  State<StatefulWidget> createState() {
    return _EditStickyPageState();
  }
}

class _EditStickyPageState extends State<EditStickyPage> {

  final _titleController = TextEditingController();
  final _contentController = TextEditingController();


  @override
  void initState() {
    super.initState();
    _titleController.value = _titleController.value.copyWith(text: widget.title);
    _contentController.value =
        _contentController.value.copyWith(text: widget.content);
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        leading: IconButton(
          icon: Icon(Icons.arrow_back, color: Colors.brown[200]),
          onPressed: () => Navigator.pop(context),
        ),
        actions: <Widget>[
          Container(
              margin: const EdgeInsets.only(left: 20),
              child: Icon(Icons.alarm, color: Colors.brown[200])),
          Container(
              margin: const EdgeInsets.only(left: 20),
              child: Icon(Icons.vertical_align_top, color: Colors.brown[200])),
          Container(
              margin: const EdgeInsets.only(left: 20, right: 15),
              child: Icon(Icons.check, color: Colors.brown[200])),
        ],
      ),
      body: Container(
        padding: const EdgeInsets.all(10),
        child: Column(
          children: <Widget>[
            TextField(
              controller: _titleController,
              decoration: InputDecoration(hintText: "标题"),
              style: TextStyle(fontSize: 20),
            ),
            Expanded(
              child: Container(
                  margin: EdgeInsets.only(top: 10),
                  child: TextField(
                    controller: _contentController,
                    decoration: InputDecoration.collapsed(hintText: "内容"),
                    maxLines: null,
                  )),
            )
          ],
        ),
      ),
    );
  }
}
