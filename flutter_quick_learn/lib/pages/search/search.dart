import 'package:flutter/material.dart';

class SearchPage extends StatelessWidget {
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
          autofocus: true,
        ),
      ),
    );
  }
}
