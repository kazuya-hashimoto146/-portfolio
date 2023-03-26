# -*- coding: utf-8 -*-

import utils
import random

print('じゃんけんスタート！')
#検証段階では文字列はクォーテーションで囲む
player_name = input('あなたの名前を教えてください：')


while True:
    print('どれを出しますか？（0: グー, 1: チョキ, 2: パー）')
    player_hand = int(input('0〜2の数字で指定してください：'))

    #0〜2の正しい数値が入っているか判定
    if utils.validate(player_hand):
        cpu_hand = random.randint(0, 2)
        
        #名前が入力されているか判定し、分岐させる
        if player_name == '':
            utils.print_hand(player_hand)
        else:
            utils.print_hand(player_hand, player_name)

        utils.print_hand(cpu_hand, 'CPU')
        
        result = utils.judge(player_hand, cpu_hand)
        if result == 'あいこ':
            print('対戦結果は' + result + 'です')
        else:
            print('対戦結果は' + result + 'です')
            break
    else:
        print('入力された値が正しくありません')
