import 'package:flutter/material.dart';

class EditStickyPage extends StatelessWidget {
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
            TextField(decoration: InputDecoration(hintText: "标题")),
            TextField(decoration: InputDecoration.collapsed(hintText: "内容"))
          ],
        ),
      ),
    );
  }
}
