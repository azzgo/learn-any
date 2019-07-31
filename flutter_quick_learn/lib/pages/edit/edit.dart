import 'package:flutter/material.dart';
import 'package:quiver/strings.dart';
import 'package:flutter_quick_learn/db/models/StickyProvider.dart';
import 'package:flutter_quick_learn/db/models/Sticky.dart';

class EditStickyPage extends StatefulWidget {
  final int id;
  final String title;
  final String content;

  EditStickyPage({this.id, this.title, this.content});

  @override
  State<StatefulWidget> createState() {
    return _EditStickyPageState();
  }
}

class _EditStickyPageState extends State<EditStickyPage> {
  TextEditingController _titleController;
  TextEditingController _contentController;

  Sticky _sticky;

  _EditStickyPageState() {
    _sticky = Sticky(title: widget.title, content: widget.content);
    _sticky.id = widget.id;
  }

  @override
  void initState() {
    super.initState();

    _titleController = TextEditingController(text: _sticky.title)
      ..addListener(() {
        _sticky.title = _titleController.text;
      });
    _contentController = TextEditingController(text: _sticky.content)
      ..addListener(() {
        _sticky.content = _contentController.text;
      });
  }

  void _saveSticky() async {
    if (!_validateData()) {
      Navigator.pop(context);
      return;
    }

    StickyProvider stickyProvider = new StickyProvider();
    await stickyProvider.open();

    if (_sticky.id == null) {
      _sticky.id = await _createSticky(stickyProvider);
    } else {
      await _updateSticky(stickyProvider);
    }

    await stickyProvider.close();
  }

  bool _validateData() {
    if (isEmpty(_titleController.text) && isEmpty(_contentController.text)) {
      return false;
    }

    return true;
  }

  Future _createSticky(StickyProvider stickyProvider) {
    return stickyProvider.insertSticky(
        Sticky(title: _titleController.text, content: _contentController.text));
  }

  Future _updateSticky(StickyProvider stickyProvider) {
    return stickyProvider.updateSticky(_sticky);
  }

  _deleteSticky() async {
    StickyProvider stickyProvider = new StickyProvider();
    await stickyProvider.open();
    if (_sticky.id != null) {
      await stickyProvider.deleteSticky(_sticky.id);
    }
    await stickyProvider.close();

    Navigator.pop(context);
    return;
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
            margin: const EdgeInsets.only(left: 20),
            child: IconButton(
                icon: Icon(Icons.delete, color: Colors.brown[200]),
                onPressed: _deleteSticky),
          ),
          Container(
              margin: const EdgeInsets.only(left: 10, right: 5),
              child: IconButton(
                  icon: Icon(Icons.check, color: Colors.brown[200]),
                  onPressed: _saveSticky)),
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
