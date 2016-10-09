Go Practice

1. sequential program
----------------------
  cd seq  
  docker build -t seq .  
  docker run -it seq /bin/bash  

*results are in the /tmp/*.txt.out*

2. Concurrent program
----------------------
  cd llel  
  docker build -t llel .  
  docker run -it llel /bin/bash  

*result file is in /tmp/results.txt
For sorted order cat /tmp/results.txt |sort*
