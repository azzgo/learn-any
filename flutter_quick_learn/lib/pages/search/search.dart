import 'package:flutter/material.dart';

import 'package:flutter_quick_learn/db/models/StickyProvider.dart';
import 'package:flutter_quick_learn/db/models/Sticky.dart';
import 'package:flutter_quick_learn/pages/edit/edit.dart';
import 'package:flutter_quick_learn/pages/home/stickyItem.dart';


class SearchPage extends StatefulWidget {
  @override
  State<StatefulWidget> createState() {
    return _SearchPageState();
  }

}

class _SearchPageState extends State<SearchPage> {
  Iterable<Sticky> stikiesIterator;

  void _searchStickies(String text) async {
    StickyProvider stickyProvider = StickyProvider();
    await stickyProvider.open();

   var stikiesIterator = await stickyProvider.fuzzyQuery(text);
   this.setState(() {
     this.stikiesIterator = stikiesIterator;
   });

   await stickyProvider.close();
  }

  void navigateToEditPage({int id, String title, String content}) async {
    await Navigator.push(
        context,
        MaterialPageRoute(
            builder: (context) => EditStickyPage(
              id: id,
              title: title,
              content: content,
            )));
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        leading: IconButton(
            icon: Icon(Icons.arrow_back),
            onPressed: () {
              Navigator.pop(context);
            }),
        title: TextField(
          decoration: InputDecoration(
              hintText: "搜索便签",
              hintStyle: TextStyle(color: Colors.brown[200]),
              suffixIcon: Icon(Icons.remove)),
          style: const TextStyle(color: Colors.white),
          onChanged: _searchStickies,
          autofocus: true,
        ),
      ),
      body: ListView.builder(itemBuilder: (context, index) {
        var sticky = stikiesIterator.toList()[index];
        return GestureDetector(
            onTap: () => navigateToEditPage(
                title: sticky.title, content: sticky.content, id: sticky.id),
            child:
            StickyItem(sticky.title, sticky.content, sticky.modifyTime));
      },
      itemCount: (stikiesIterator ?? []).length,),
    );
  }


}
