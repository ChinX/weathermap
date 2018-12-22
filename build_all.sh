#!/usr/bin/env bash
CUR_DIR=`pwd`
command="$1"
if [ "${command}" == "base" ]; then
    cd ${CUR_DIR}/base-images
     #make all
fi
for item in forecast fusionweather weather weather-beta weathermapweb; do
    cd ${CUR_DIR}/${item}
    make image
done
cd ${CUR_DIR}