{{$playercount := (toInt ((dbGet 204255221017214977 "games_count").Value))}}
{{$slice := (dbGet 204255221017214977 "players").Value}}
{{$plural := ""}}
 
{{if gt $playercount 1}}
	{{$plural = "s"}}
{{end}}
 
{{/* OPTIONAL: Put a channel ID that logs activity from YAGPDB */}}
{{sendMessage YAG_LOG_CHANNEL_ID (joinStr "" "Removing roles from " $playercount " player" $plural "...")}}
 
{{range $i := (seq 1 (add $playercount 1))}}
	{{takeRoleID (index $slice $i) ACTIVE_PLAYER_ROLE_ID}}
	{{takeRoleID (index $slice $i) LOSING_PLAYER_ROLE_ID}}
	{{sendMessage nil $i}}
{{end}}
 
{{sendMessage YAG_LOG_CHANNEL_ID "Completed"}}
