
import 'package:flutter/material.dart';

class StickyItem extends StatelessWidget {
  final String title;
  final String content;
  final String modifyTime;

  StickyItem(this.title, this.content, this.modifyTime);

  @override
  Widget build(BuildContext context) {


    return Card(
      child: Container(
        height: 130,
        padding: EdgeInsets.all(10),
        child: Column(
            mainAxisSize: MainAxisSize.max,
            crossAxisAlignment: CrossAxisAlignment.start,
            children: <Widget>[
              Text(title, style: TextStyle(fontSize:  20)),
              Expanded(child: Text(content)),
              Text(modifyTime)
            ],
        )
      ),
    );
  }

}