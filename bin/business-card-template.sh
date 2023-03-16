#!/bin/bash

# Output the business card to a file
cat <<EOF > ./public/businessCard
┌──────────────────────────────────────────────────────────┐
│                                                          │
│                       Marcus Price                       │
│                    Software Developer                    │
│                                                          │
│                                                          │
│       email: marcusprice88@gmail.com                     │
│      mobile: 206.280.3302                                │
│         web: https://marcusprice.me                      │
│      github: https://github.com/marcusprice              │
│    linkedin: https://linkedin.com/in/marcus-price        │
│                                                          │
│        card: curl https://marcusprice.me                 │
│                                                          │
└──────────────────────────────────────────────────────────┘
EOF

# Print a message indicating that the file was created
echo "business card was created, here is the resulting output:"
cat public/businessCard
