{{$lobby := dbGet 204255221017214977 "games_lobby"}}
{{$count := dbGet 204255221017214977 "games_count"}}
 
{{if not $lobby.Value}}
	{{sendMessage nil "Lobby is closed!"}}
{{else if eq (toInt $count.Value) 10}}
	{{sendMessage nil "Lobby is full!"}}
{{else if (hasRoleID ACTIVE_PLAYER_ROLE_ID)}}
	{{sendMessage nil "You're already in the game!"}}
{{else}}
	{{dbSet 204255221017214977 "games_count" (add (toInt $count.Value) 1)}}
	{{sendMessage nil (joinStr "" .User.Mention " has entered the lobby! " (add (toInt $count.Value) 1) "/10 players")}}
	{{addRoleID ACTIVE_PLAYER_ROLE_ID}}
 
{{$slice := (dbGet 204255221017214977 "players").Value}}
{{$cslice_converted := (cslice).AppendSlice $slice}}
{{dbSet 204255221017214977  "players" ($cslice_converted.AppendSlice (cslice (toString .User.ID)))}}
 
{{end}}
