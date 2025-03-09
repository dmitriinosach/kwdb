import http from 'k6/http';
import { sleep, check } from 'k6';

export const options = {
  vus: 10,
  duration: '30s',
  cloud: {
    projectID: 3753077,
    // Test runs with the same name groups test runs together
    name: 'YOUR TEST NAME'
  },
};

export default function() {

  let val = "hfouwhepori jfdsfjk as[dpifh [pkmfsda[pf kas]df k-20u-439jksdof sa['df sad]f [okasd]f -0idsa-f s0ak9df]- as0dkf= a-sdkf=as9 if2-1ok3ef,d psdfjp[sadjf"
  let vals = "12313t dcvs ga"
  let res = http.get('http://188.242.160.102:713/?key='+getRandomInt(200000)+"&value="+vals);

  check(res, { "status is 200": (res) => res.status === 200 });

  sleep(1);


}

function getRandomInt(max) {
  return Math.floor(Math.random() * max);
}