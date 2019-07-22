import 'Item.dart';
import 'ShoppingCard.dart';

void main() {
  ShoppingCard sc = ShoppingCard('张三', '123456');
  sc.bookings = [Item('青藏高原之谜', 10.0), Item('外星人来访记录', 50.84)];
  print(sc.getInfo());
}