{{/* SET THE TRIGGER TYPE TO REACTION, ADDED REACTIONS ONLY */}}
{{/* ONLY RUN IN THE VERIFICATION CHANNEL, NO ROLE RESTRICTIONS */}}

{{/* Used to verify users with a "secret" emoji, such as those given in server rules. */}}
{{/* It will clear any emoji as fast as it can, right or wrong. *}}

{{/* Replace with emoji of choice, use text for custom emojis */}}
{{if eq .Reaction.Emoji.Name "👍"}}
    {{removeRoleID UNVERIFIED_ROLE_ID 0}}
    
    {{/* Optional: Depending on how you wish to manage permissions for @ everyone. */}}
    {{addRoleID VERIFIED_ROLE_ID}}
{{end}}
 
{{deleteAllMessageReactions VERIFICATION_CHANNEL_ID MESSAGE_WITH_REACTIONS_ID}}
