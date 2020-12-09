{{$count := dbGet 204255221017214977 "games_count"}}
 
{{$timer := randInt 10 25}} {{/* {{sendMessage nil $timer}} Uncomment to send a message every passing second, debug */}}
{{dbSet 204255221017214977  "timer" $timer}}
{{/* execCC ## of fuse */}}
{{execCC ## nil 0 "nada"}}
 
{{$randplayer := randInt 1 (toInt $count.Value)}}
{{/* Just making the first player to join have bomb and seeing how that goes */}}
 
{{$slice := (dbGet 204255221017214977 "players").Value}}
{{$bomberman := index $slice 1}}
	{{if (targetHasRoleID $bomberman ACTIVE_PLAYER_ROLE_ID)}}
		{{dbSet 204255221017214977 "bomberman" $bomberman}}
		{{sendMessage nil (joinStr "" "<@" $bomberman ">" ", you have the bomb! Type `-p @Someone` to get rid of it!")}}
	{{else}}
		This person isn't properly licensed!
{{end}}
