# TripsToMars
 Development of a search engine for trips to Mars

Программа для покупки путёвок на Марс - 
генерирует 10 случайных билетов и отображает их в табличном виде.

В коде используются переменные, константы, switch, if и for. Для отображения, выравнивания текста и генерации случайных чисел задействованы пакеты fmt и math/rand.

В таблице четыре столбца:

- Космическая станция (Spaceline), что предоставляет услуги;
- Продолжительность (Days) в днях поездки на Марс в один конец;
- Покрывает ли цена поездку туда и обратно (Trip type);
- Цена (Price) в миллионах долларов.

Для каждого билета случайным образом выбирается космическая станция: Space Adventures, SpaceX или Virgin Galactic.

Датой отправления на каждом билете значится 13 Октября 2020 года. В этот день Марс находится на расстоянии 62 100 000 км от Земли.

Скорость космического корабля выбрана случайным образом из диапазона от 30 до 60 км/c. 
Это определяет продолжительность поездки на Марс, а также цену билета. Более быстрые корабли намного дороже. 
Цены на билеты варьируются от $50 до $80 миллионов. Цена для поездки туда-обратно удваивается.
 
