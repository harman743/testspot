# Luke Phillips - J00624776

# Program is a game of rock, paper, scissors
# Player chooses rock, paper, or scissors
# Program displays game result (win, lose, draw)
# Program then asks if they want to play again

.data
	choose: .asciiz "Enter [1] to choose rock, enter [2] to choose paper, enter [3] to choose scissors: "
	players_choice: .asciiz "You chose "
	computers_choice: .asciiz "The computer chose "
	rock: .asciiz "rock"
	paper: .asciiz "paper"
	scissors: .asciiz "scissors"
	period: .asciiz ".\n"
	win: .asciiz "You won!"
	lose: .asciiz "You lose."
	draw: .asciiz "The game was a draw."
	again: .asciiz "Enter [1] to play again or enter [0] to quit: "
	nl: .asciiz "\n"
	
.text

main:
	jal player_choose
	move $s1, $v0	# store player's choice in $s1
	
	jal computer_choose
	move $s2, $v0	# store computer's choice in $s2
	
	jal get_result
	move $s3, $v0	# store game result in $s3
	
	jal output_result
	
	# print new line
	li $v0, 4
	la $a0, nl
	syscall
	
	# ask to play again
	li $v0, 4
	la $a0, again
	syscall
	
	# read integer
	li $v0, 5
	syscall
	
	# if input is 1 then loop, else end program
	beq $v0, 1, main
	bne $v0, 1, end_program
    
# ---------------------------------------------------------------------------
# Subroutine player_choose
#
# No arguements
#
# Return value in $v0: player's choice (1 for rock, 2 for paper, 3 for scissors)
# ---------------------------------------------------------------------------
player_choose:
	
	li $v0, 4	# print string
	la $a0, choose	# ask user to choose rock, paper, or scissors
	syscall
	
	li $v0, 5	# read integer
	syscall
	
	beq $v0, 1, valid_choice
	beq $v0, 2, valid_choice
	beq $v0, 3, valid_choice
	
	j player_choose	# choice was invalid, ask for input again
	
valid_choice:	
	jr $ra
	

# ---------------------------------------------------------------------------
# Subroutine computer_choose
#
# No arguements
#
# Return value in $v0: computer's choice (1 for rock, 2 for paper, 3 for scissors)
# ---------------------------------------------------------------------------
computer_choose:
	
	li $a1, 3  # set the upper bound of the integer to be generated
	li $v0, 42  # System call code 42 (random integer in range)
	syscall
	add $a0, $a0, 1  # set the lower bound of the integer that is generated
	move $v0, $a0	# store computer's choice in $v0
	
	jr $ra


# ---------------------------------------------------------------------------
# Subroutine get_result
#
# Argument(s): Player's choice in $s1, computer's choice in $s2
#
# Return value in $v0: 1 if player wins, 2 if player loses, 3 if it's a draw
# ---------------------------------------------------------------------------
get_result:
	beq $s1, 1, chose_rock
	beq $s1, 2, chose_paper
	beq $s1, 3, chose_scissors
	
	chose_rock:
		# if CPU chose rock then it's a draw
		beq $s2, 1, game_draw
		# if CPU chose paper then player loses
		beq $s2, 2, game_lose
		# if CPU chose scissors then player wins
		beq $s2, 3, game_win
	chose_paper:
		# if CPU chose rock then player wins
		beq $s2, 1, game_win
		# if CPU chose paper then it's a draw
		beq $s2, 2, game_draw
		# if CPU chose scissors then player loses
		beq $s2, 3, game_lose
	chose_scissors:
		# if CPU chose rock then player loses
		beq $s2, 1, game_lose
		# if CPU chose paper then player wins
		beq $s2, 2, game_win
		# if CPU chose scissors then it's a draw
		beq $s2, 3, game_draw
		
	game_win:
		# set $v0 to 1
		li $v0, 1
		jr $ra
	game_lose:
		# set $v0 to 2
		li $v0, 2
		jr $ra
	game_draw:
		# set $v0 to 3
		li $v0, 3
		jr $ra

# ---------------------------------------------------------------------------
# Subroutine output_result
#
# Argument(s): Game result in $s3
#
# No return value
# ---------------------------------------------------------------------------
output_result:
	# print new line
	li $v0, 4
	la $a0, nl
	syscall
	
	# print "You chose "
	li $v0, 4
	la $a0, players_choice
	syscall
	
	# push $ra to stack
	addi $sp,$sp,-4
	sw $ra,0($sp)
	
	# move player's choice value into $a1
	move $a1, $s1
	# evaluate $a1 to get the correct string ("rock", "paper", or "scissors")
	jal choice_string	
	
	# pop $ra off stack
	lw $ra,0($sp)
	add $sp,$sp,4
	
	# print the string corresponding to the player's choice
	li $v0, 4
	syscall
	
	# print period
	li $v0, 4
	la $a0, period
	syscall
	
	# print "The computer chose "
	li $v0, 4
	la $a0, computers_choice
	syscall
	
	# push $ra to stack
	addi $sp,$sp,-4
	sw $ra,0($sp)
	
	# move computer's choice value into $a1
	move $a1, $s2
	# evaluate $a1 to get the correct string ("rock", "paper", or "scissors")
	jal choice_string	
	
	# pop $ra off stack
	lw $ra,0($sp)
	add $sp,$sp,4
	
	# print the string corresponding to the computer's choice
	li $v0, 4
	syscall
	
	# print period
	li $v0, 4
	la $a0, period
	syscall
	
	# push $ra to stack
	addi $sp,$sp,-4
	sw $ra,0($sp)
	
	# evaluate $s3 to get the correct string ([win], [lose], or [draw])
	jal result_string
	
	# pop $ra off stack
	lw $ra,0($sp)
	add $sp,$sp,4
	
	# print the string corresponding to the game result value
	li $v0, 4
	syscall
	
	# print new line
	li $v0, 4
	la $a0, nl
	syscall
	
	jr $ra
	
# ---------------------------------------------------------------------------
# Subroutine choice_string
#
# Argument(s): choice value in $a1
#
# Return value in $a0: address of string "rock" or "paper" or "scissors"
# ---------------------------------------------------------------------------
choice_string:
	beq $a1, 1, say_rock
	beq $a1, 2, say_paper
	beq $a1, 3, say_scissors
	
	say_rock:
		# store the address of the string "rock" in $a0
		la $a0, rock
		jr $ra
	say_paper:
		# store the address of the string "paper" in $a0
		la $a0, paper
		jr $ra
	say_scissors:
		# store the address of the string "scissors" in $a0
		la $a0, scissors
		jr $ra
	
# ---------------------------------------------------------------------------
# Subroutine result_string
#
# Argument(s): game result value in $s3
#
# Return value in $a0: address of string [win],[lose], or [draw]
# ---------------------------------------------------------------------------
result_string:
	beq $s3, 1, say_win
	beq $s3, 2, say_lose
	beq $s3, 3, say_draw
	
	say_win:
		# store the address of the string [win] in $a0
		la $a0, win
		jr $ra
	say_lose:
		# store the address of the string [lose] in $a0
		la $a0, lose
		jr $ra
	say_draw:
		# store the address of the string [draw] in $a0
		la $a0, draw
		jr $ra


end_program:
