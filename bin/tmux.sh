#!/bin/bash
# Session Name
session="server"

# Start New Session with our name
tmux new-session -d -s $session

tmux rename-window -t 0 "code"
tmux send-keys -t "code" "zsh" C-m "clear" C-m "vim ." C-m

tmux new-window
tmux rename-window -t 1 "git"
tmux send-keys -t "git" "clear" C-m "git status" C-m

tmux new-window
tmux rename-window -t 2 "shell"

tmux new-window
tmux rename-window -t 3 "dev-server"
tmux send-keys -t "dev-server" "make dev" C-m

tmux new-window
tmux rename-window -t 4 "curl"

tmux attach-session -t $SESSION
