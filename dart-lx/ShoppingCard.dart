import 'Item.dart';
import 'Meta.dart';
import 'dateTime.util.dart' as utils;

class ShoppingCard extends Meta {
  DateTime dateTime;
  String code;

  List<Item> bookings;

  ShoppingCard(name, this.code): dateTime = DateTime.now(), super(name, 0);

  double get price {
    double sum = 0.0;

    for (var i in bookings) {
      sum + i.price;
    }

    return sum;
  }

  getInfo() {
    return '''购物车信息:
------------------------
用户名: $name
优惠码: $code
总价:   ${price.toString()}
日期:   ${utils.format(dateTime, "YYYY-MM-DD HH:mm:ss")}
------------------------
    ''';
  }
}
