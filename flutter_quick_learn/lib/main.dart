import 'package:flutter/material.dart';
import 'package:flutter_quick_learn/pages/home/home.dart';

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: '便签',
      theme: ThemeData(
        primarySwatch: Colors.brown,
      ),
      home: HomePage(title: '便签'),
    );
  }
}

