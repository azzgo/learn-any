
import 'package:flutter/material.dart';
import 'package:flutter_quick_learn/db/models/StickyProvider.dart';
import 'package:flutter_quick_learn/db/models/Sticky.dart';
import 'stickyItem.dart';
import 'dart:developer' as developer;

class HomePage extends StatefulWidget {
  HomePage({Key key, this.title}) : super(key: key);

  final String title;

  @override
  _HomePageState createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {

  Iterable<Sticky> stickies;

  @override
  void initState() {
    super.initState();
    initStickiesData();
  }

  void initStickiesData() async {
    var stickyProvider =  StickyProvider();
    await stickyProvider.open();
    var stickies = await stickyProvider.getStickies();
    this.setState(() {
      developer.log(stickies.toString());
      this.stickies = stickies;
    });
  }

  @override
  Widget build(BuildContext context) {


    return Scaffold(
      appBar: AppBar(
        title: Text(widget.title),
      ),
      body: ListView(
        children: stickies != null ? stickies.map((sticky) => StickyItem(sticky.title, sticky.content, sticky.modifyTime.toString())).toList() : []
      ),
      floatingActionButton: FloatingActionButton(
        backgroundColor: Colors.brown,
        tooltip: 'Increment',
        child: Icon(Icons.add),
      ),
    );
  }
}