d=$PWD
# Absolute path to this script, e.g. /home/user/bin/foo.sh
SCRIPT=$(readlink -f "$0")
# Absolute path this script is in, thus /home/user/bin
SCRIPTPATH=$(dirname "$SCRIPT")
echo $SCRIPTPATH
cd $SCRIPTPATH
echo "Start elasticsearch.."
cd elasticsearch-5.6.1
bin/elasticsearch -d
echo "Start kibana.."
cd $SCRIPTPATH
cd kibana-5.6.1-linux-x86_64
nohup bin/kibana &
echo "Start tdcs_redirect.."
cd $SCRIPTPATH
cd tdcs_redirect
nohup ./tdcs_redirect &
echo "Start logstash.."
cd $SCRIPTPATH
cd logstash-5.6.1
nohup bin/logstash -f ../etag/m03a.conf --path.data data.m03a > logs/m03a.log &
nohup bin/logstash -f ../etag/m05a.conf --path.data data.m05a > logs/m05a.log &
nohup bin/logstash -f ../etag/rain.conf --path.data data.rain > logs/rain.log &
cd $d
