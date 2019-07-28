import 'package:flutter/material.dart';
import 'package:flutter_quick_learn/db/models/StickyProvider.dart';
import 'package:flutter_quick_learn/db/models/Sticky.dart';
import 'package:flutter_quick_learn/pages/edit/edit.dart';
import 'package:flutter_quick_learn/pages/search/search.dart';
import 'stickyItem.dart';

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
    var stickyProvider = StickyProvider();
    await stickyProvider.open();
    var stickies = await stickyProvider.getStickies();
    this.setState(() {
      this.stickies = stickies;
    });
    stickyProvider.close();
  }

  void navigateToEditPage({int id, String title, String content}) {
    Navigator.push(
        context,
        MaterialPageRoute(
            builder: (context) => EditStickyPage(
                  id: id,
                  title: title,
                  content: content,
                )));
  }

  void navigateToSearchPage() {
    Navigator.push(
        context, MaterialPageRoute(builder: (context) => SearchPage()));
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(widget.title),
        actions: <Widget>[
          IconButton(
              icon: Icon(Icons.search, color: Colors.brown[200]),
              onPressed: navigateToSearchPage),
        ],
      ),
      body: ListView.builder(
        itemBuilder: (context, index) {
          var sticky = stickies.toList()[index];
          return GestureDetector(
              onTap: () => navigateToEditPage(
                  title: sticky.title, content: sticky.content, id: sticky.id),
              child:
                  StickyItem(sticky.title, sticky.content, sticky.modifyTime));
        },
        itemExtent: 130,
        itemCount: stickies.length,
      ),
      floatingActionButton: FloatingActionButton(
        backgroundColor: Colors.brown,
        tooltip: 'Increment',
        child: Icon(Icons.add),
        onPressed: navigateToEditPage,
      ),
    );
  }
}
