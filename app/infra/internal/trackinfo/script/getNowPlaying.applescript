set appName to ""
set trackName to ""
set albumName to ""
set artistName to ""
set trackDuration to 0
set playerPosition to 0

tell application "System Events"
  if (exists process "Music") then
    tell application "Music"
      if player state is playing then
        set appName to "Music"
        set trackName to name of current track
        set albumName to album of current track
        set artistName to artist of current track
        set trackDuration to duration of current track
        set playerPosition to player position
      end if
    end tell
  end if

  if (exists process "Audirvana Origin") then
    tell application "Audirvana Origin"
      if player state is Playing then
        set appName to "Audirvana Origin"
        set trackName to playing track title
        set albumName to playing track album
        set artistName to playing track artist
        set trackDuration to playing track duration
        set playerPosition to player position
      end if
    end tell
  end if
end tell

if appName = "" then
  return
end if

set jsonResult to "{"
set jsonResult to jsonResult & "\"appName\": \"" & appName & "\", "
set jsonResult to jsonResult & "\"track\": \"" & trackName & "\", "
set jsonResult to jsonResult & "\"album\": \"" & albumName & "\", "
set jsonResult to jsonResult & "\"artist\": \"" & artistName & "\", "
set jsonResult to jsonResult & "\"duration\": " & trackDuration & ", "
set jsonResult to jsonResult & "\"position\": " & playerPosition
set jsonResult to jsonResult & "}"

return jsonResult
