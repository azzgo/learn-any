import 'package:flutter/material.dart';
import 'package:flutter_quick_learn/db/models/StickyProvider.dart';
import 'package:flutter_quick_learn/db/models/Sticky.dart';
import 'package:flutter_quick_learn/pages/edit/edit.dart';
import 'package:flutter_quick_learn/pages/search/search.dart';
import 'models/StickyWithChecked.dart';
import 'stickyItem.dart';

class HomePage extends StatefulWidget {
  HomePage({Key key, this.title}) : super(key: key);

  final String title;

  @override
  _HomePageState createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> with WidgetsBindingObserver {
  List<Sticky> _stickies;

  List<StickyWithChecked> _mupltipleStickies;

  bool _isMutipleMode = false;

  bool get isMultipleMode {
    return this._isMutipleMode;
  }

  void set isMultipleMode(bool val) {
    if (val && !this._isMutipleMode) {
      _mupltipleStickies = _stickies
          .map((item) => StickyWithChecked(sticky: item, checked: false))
          .toList();
    }

    if (!val) {
      _mupltipleStickies = null;
    }
    this._isMutipleMode = val;
  }

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
      this._stickies = stickies.toList();
    });
    stickyProvider.close();
  }

  void navigateToEditPage({int id}) async {
    await Navigator.push(context,
        MaterialPageRoute(builder: (context) => EditStickyPage(id: id)));
    this.initStickiesData();
  }

  void navigateToSearchPage() {
    Navigator.push(
        context, MaterialPageRoute(builder: (context) => SearchPage()));
  }

  void _enterMultipleMode() {
    setState(() {
      isMultipleMode = true;
    });
  }

  AppBar _getAppBar() {
    return AppBar(
      title: Text(widget.title),
      actions: <Widget>[
        IconButton(
            icon: Icon(Icons.search, color: Colors.brown[200]),
            onPressed: navigateToSearchPage),
      ],
    );
  }

  AppBar _getCheckableListViewAppBar() {
    return AppBar(
      leading: IconButton(
          icon: Icon(Icons.close),
          onPressed: () {
            setState(() {
              isMultipleMode = false;
            });
          }),
      actions: <Widget>[
        IconButton(
          icon: Icon(Icons.delete),
          onPressed: () async {
            StickyProvider stickyProvider = StickyProvider();
            await stickyProvider.open();
            _mupltipleStickies.where((item) => item.checked).forEach((item) {
              stickyProvider.deleteSticky(item.sticky.id);
            });

            await stickyProvider.close();

            this.isMultipleMode = false;

            this.initStickiesData();
          },
        ),
        IconButton(
            icon: Icon(
              Icons.check_box,
              color: Colors.white,
            ),
            onPressed: () {
              if (_mupltipleStickies.every((item) => item.checked)) {
                _mupltipleStickies.forEach((item) {
                  item.checked = false;
                });
              } else {
                _mupltipleStickies.forEach((item) {
                  item.checked = true;
                });
              }
              this.setState(() {});
            })
      ],
    );
  }

  ListView _getListView() {
    return ListView.builder(
      itemBuilder: (context, index) {
        var sticky = _stickies[index];
        return GestureDetector(
            onTap: () => navigateToEditPage(id: sticky.id),
            onLongPress: _enterMultipleMode,
            child: StickyItem(sticky.title, sticky.content, sticky.modifyTime));
      },
      itemExtent: 130,
      itemCount: _stickies?.length ?? 0,
    );
  }

  ListView _getCheckableListView() {
    return ListView.builder(
      itemBuilder: (context, index) {
        var stickyWithChecked = _mupltipleStickies[index];
        var sticky = stickyWithChecked.sticky;
        return GestureDetector(
            onTap: () => {
                  setState(() {
                    stickyWithChecked.checked = !stickyWithChecked.checked;
                  })
                },
            child: Row(children: <Widget>[
              Icon(stickyWithChecked.checked
                  ? Icons.check_box
                  : Icons.check_box_outline_blank),
              Expanded(
                child:
                    StickyItem(sticky.title, sticky.content, sticky.modifyTime),
              )
            ]));
      },
      itemExtent: 130,
      itemCount: _mupltipleStickies?.length ?? 0,
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: isMultipleMode? _getCheckableListViewAppBar() : _getAppBar(),
      body: isMultipleMode ? _getCheckableListView() : _getListView(),
      floatingActionButton: !isMultipleMode ? FloatingActionButton(
        backgroundColor: Colors.brown,
        tooltip: '增加便签',
        child: Icon(Icons.add),
        onPressed: navigateToEditPage,
      ) : null,
    );
  }
}
