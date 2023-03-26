# -*- coding: utf-8 -*-

def validate(hand):
    if hand < 0 or hand > 2:
        return False
    return True

def print_hand(hand, name = 'プレイヤー'):
    hands = ['グー', 'チョキ', 'パー']
    print(name + 'さんは' + hands[hand] + 'を出しました！')

def judge(player, cpu):
    if player == cpu:
        return 'あいこ'
    elif player == 0 and cpu == 1:
        return '勝ち'
    elif player == 1 and cpu == 2:
        return '勝ち'
    elif player == 2 and cpu == 0:
        return '勝ち'
    else:
        return '負け'