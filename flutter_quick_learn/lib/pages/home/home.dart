
import 'package:flutter/material.dart';
import 'package:flutter_quick_learn/mocks/stickList.dart';
import 'stickyItem.dart';

class HomePage extends StatefulWidget {
  HomePage({Key key, this.title}) : super(key: key);

  final String title;

  @override
  _HomePageState createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {


  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(widget.title),
      ),
      body: ListView(
        children: STICKY_LIST_DATA.map((item) => StickyItem(item['title'], item['content'], item['modifyTime'])).toList()
      ),
      floatingActionButton: FloatingActionButton(
        backgroundColor: Colors.brown,
        tooltip: 'Increment',
        child: Icon(Icons.add),
      ),
    );
  }
}