# -*- coding: utf-8 -*-

import utils
import random

print('じゃんけんをはじめます！')
#検証段階では文字列はクォーテーションで囲む
player_name = input('名前を入力してください：')

print('何を出しますか？（0: グー, 1: チョキ, 2: パー）')
player_hand = int(input('数字で入力してください：'))

#0〜2の正しい数値が入っているか判定
if utils.validate(player_hand):
    computer_hand = random.randint(0, 2)
    
    #名前が入力されているか判定し、分岐させる
    if player_name == '':
        utils.print_hand(player_hand)
    else:
        utils.print_hand(player_hand, player_name)

    utils.print_hand(computer_hand, 'コンピューター')
    
    result = utils.judge(player_hand, computer_hand)
    print('結果は' + result + 'でした')
else:
    print('正しい数値を入力してください')
