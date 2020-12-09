{{/* THIS GAME IS INSPIRED BY SUBMER#0001, OWNER OF THE PREDECESSORS SERVER ON DISCORD, THIS IS ONLY MY ATTEMPT TO REPLICATE IT */}}

{{$args := parseArgs 0 ""  (carg "int" "seconds to join lobby") (carg "string" "name of game")}}
{{$time := $args.Get 0}} {{$game := $args.Get 1}} 

{{/* All dbSet will use YAGPDB's ID */}}
{{dbSet 204255221017214977 "games_lobby" 1}}
{{dbSet 204255221017214977  "games_status" 0}}
{{dbSet 204255221017214977 "games_count" 0}}
{{dbSet 204255221017214977  "players" (cslice "Players:")}}
 
{{/* Default to "Bomb Time" if no other game name is given */}}
{{if not $game}} {{$game = "Bomb Time"}} {{end}}
{{if not $time}}
	{{$time = 30}}
{{else if lt $time 10}}
	{{$time = 10}}
{{else if gt $time 60}}
	{{$time = 60}}
{{end}}
 
{{sendMessage nil (joinStr " " "A game of" $game "will start! Join the lobby with this command: `-join lobby`")}}
 
{{$id := sendMessageRetID nil (joinStr " " $time "Seconds Remaining...")}} {{$time = sub $time 1}} {{sleep 1}}
{{range $i := seq 0 ($time) }}
	{{editMessage nil $id (joinStr " " $time "Seconds Remaining...")}} {{$time = sub $time 1}} {{sleep 1}}
		{{if eq $time 0}}
			{{$count := toInt (dbGet 204255221017214977 "games_count").Value}}
				{{if (lt $count 2)}}
					{{dbSet 204255221017214977 "games_lobby" 0}}
					{{dbSet 204255221017214977  "games_status" 0}}
					{{deleteMessage nil $id 0}}
					{{sendMessage nil "Not enough players :("}}
          				{{/* CC ## to clear roles that were initially given to players in lobby */}} 
					{{execCC ## nil 0 0}}
				{{else}}
					{{dbSet 204255221017214977 "games_lobby" 0}}
          				{{/* SET THE CHANNEL ID YOU WANT THE GAME TO TAKE PLACE IN */}}
					{{editMessage nil $id "Game has begun! Play at <#GAME_CHANNEL_ID>"}}
					{{sendMessage nil "Game has begun! Play at <#GAME_CHANNEL_ID>"}}
					{{dbSet 204255221017214977  "games_status" 1}}
					{{dbSet 204255221017214977  "games_alive" $count}}
          				{{/* CC ## for Bombspawn */}}
					{{execCC ## GAME_CHANNEL_ID 5 "nada"}}
				{{end}}
		{{end}}
{{end}}
