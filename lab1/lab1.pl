% База Знаний на основе информации об игроках и командах в ксго

% Ввод базы игроков
player(zywoo).
player(simple).
player(shiro).
player(deko).
player(monesy).
player(smooya).
player(dexter).
player(niko).
player(xantares).
player(ropz).
player(device).
player(jame).

% Ввод базы команд
team(spirit).
team(vitality).
team(mibr).
team(virtuspro).
team(g2).
team(heroic).
team(astralis).
team(navi).

% Ввод страны для игроков
country(zywoo, france).
country(simple, ukrain).
country(shiro, russia).
country(monesy, russia).
country(deko, russia).
country(device, denmark).
country(jame, russia).

% Ввод стоимости игроков
player_price(zywoo, 99999999999).
player_price(simple, 9).
player_price(shiro, 100022).
player_price(monesy, 321099999999894).
player_price(deko, 231490).
player_price(device, 213).
player_price(jame, 132).

% Максимальная стоимость игрока
max_player_price(Player, MaxPrice) :-
    player_price(Player, Price),
    \+ (player_price(_, OtherPrice), OtherPrice > Price),
    MaxPrice is Price.

% Проверка игрока с максимальной стоимостью
find_max_player(Player, MaxPrice) :-
    max_player_price(Player, MaxPrice),
    write('Игрок с максимальной стоимостью: '), write(Player), nl,
    write('Стоимость: '), write(MaxPrice).

% Условие вступления в команду нави
in_team_navi(Player):-
    player(Player),
    country(Player, ukrain).
% Условие вступления в команду спирит
in_team_spirit(Player):-
    player(Player),
    country(Player, russia).
% Проверка на адекватность цены игрока
overprice(Player):-
    player(Player),
    (player_price(Player, X), X > 4000, write("Он не стоит больше 4000 рублей");
    player_price(Player, X), X =< 4000, write("ЧЕЛ ХАРОШ МЕГАХАРОШ СУПЕРМЕГАХАРОШ")).
% Проверка на уровень игры команды
team_state(Team):-
    team(Team), (Team = spirit; Team = virtuspro),
    write("ЖËСТКАЯ ТИМА");
    write("РАКИ").
% Проверка на принадлежность игрока к сущности лягушки
is_frog(Player):-
    country(Player, X), X = france, write("Лягуха");
    write("Не лягуха").

/*
team(g2).
player(aboba).

team(spirit), player(jame).
country(jame, russia); team(aboba).

country(X, russia).
player_price(monesy, X).

team_state(spirit).
overprice(zywoo).
*/
