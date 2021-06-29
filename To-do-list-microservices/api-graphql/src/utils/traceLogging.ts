
import Jaeger from 'jaeger-client';

const ingestUrl = 'http://localhost:14268/api/traces'

export const initTracer = (serviceName: string) => {
  const config = {
    serviceName: serviceName,
    sampler: {
      type: "const",
      param: 1,
    },
    reporter: {
      logSpans: true,
      collectorEndpoint: ingestUrl,
    },
  };
  const options = {
    logger: {
      info(msg: string) {
        console.log("INFO ", msg);
      },
      error(msg: string) {
        console.log("ERROR", msg);
      },
    },
  };
  return Jaeger.initTracer(config, options);
};

