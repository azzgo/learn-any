import 'package:flutter/material.dart';
import 'package:flutter_quick_learn/db/models/StickyProvider.dart';
import 'package:flutter_quick_learn/db/models/Sticky.dart';

class EditStickyPage extends StatefulWidget {
  final int id;

  EditStickyPage({this.id});

  @override
  State<StatefulWidget> createState() {
    return _EditStickyPageState();
  }
}

class _EditStickyPageState extends State<EditStickyPage> {
  TextEditingController _titleController;
  TextEditingController _contentController;

  Sticky _sticky = Sticky();

  bool _isDirty = false;

  @override
  void initState() {
    super.initState();

    this._initStiky();
  }

  void _initStiky() async {
    StickyProvider stickyProvider = StickyProvider();
    await stickyProvider.open();

    if (widget.id != null) {
      _sticky = await stickyProvider.getSticky(widget.id);
    }

    await stickyProvider.close();

    _titleController = TextEditingController(text: _sticky.title)
      ..addListener(() {
        _isDirty = true;
        _sticky.title = _titleController.text;
      });
    _contentController = TextEditingController(text: _sticky.content)
      ..addListener(() {
        _isDirty = true;
        _sticky.content = _contentController.text;
      });

    // 因为 setState 的回调函数，不能声明为 async, 需要异步工作做完后，同步调用
    this.setState(() => null);
  }

  void _saveSticky() async {
    if (!_isDirty) {
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

    Navigator.pop(context);
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

  void _toggleFixTop() {
    setState(() {
      _sticky.isTop = !_sticky.isTop;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        leading: IconButton(
          icon: Icon(Icons.arrow_back, color: Colors.brown[200]),
          onPressed: _saveSticky,
        ),
        actions: <Widget>[
          Container(
              margin: const EdgeInsets.only(left: 20),
              child: Icon(Icons.alarm, color: Colors.brown[200])),
          Container(
              margin: const EdgeInsets.only(left: 15),
              child: IconButton(
                  icon: Icon(Icons.vertical_align_top,
                      color: _sticky.isTop ? Colors.white : Colors.brown[200]),
                  onPressed: _toggleFixTop)),
          Container(
            margin: const EdgeInsets.only(left: 5, right: 5),
            child: IconButton(
                icon: Icon(Icons.delete, color: Colors.brown[200]),
                onPressed: _deleteSticky),
          )
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
