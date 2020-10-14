# Kill all lingering related processes
function cleanup {
	pkill go-simulator
}

cd ..
exec -a go-simulator ./go-simulator &

trap cleanup EXIT

while : ; do sleep 1 ; done