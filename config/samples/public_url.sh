while true
do
  public_url=$(kubectl exec $(kubectl get pods -l=app=ngrok -o=jsonpath='{.items[0].metadata.name}') -- curl --silent  http://localhost:4040/api/tunnels | sed -nE 's/.*public_url":"https:..([^\"]*).*/\1/p')
  if [ -n $public_url ]
    then
    echo "https://"$public_url
    break
  fi
done