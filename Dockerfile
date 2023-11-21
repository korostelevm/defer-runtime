FROM public.ecr.aws/lambda/nodejs:20


COPY ./lambda/* ./lambda/
CMD [ "./lambda/index.lambdaHandler"]
