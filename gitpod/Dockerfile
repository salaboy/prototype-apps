FROM gitpod/workspace-full

USER root

#RUN sudo apt-get install apt-transport-https ca-certificates gnupg -y
#RUN echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] http://packages.cloud.google.com/apt cloud-sdk main" | tee -a /etc/apt/sources.list.d/google-cloud-sdk.list && curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key --keyring /usr/share/keyrings/cloud.google.gpg  add - && apt-get update -y && apt-get install google-cloud-cli -y
#RUN sudo apt-get install -y kubectl 
#RUN sudo apt-get install google-cloud-sdk-gke-gcloud-auth-plugin   

RUN mkdir -p /home/gitpod/kubectl/bin/
RUN curl -Lo /home/gitpod/kubectl/bin/kubectl "https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl"
RUN chmod +x /home/gitpod/kubectl/bin/kubectl
ENV PATH $PATH:/home/gitpod/kubectl/bin

# gcloud
RUN echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] http://packages.cloud.google.com/apt cloud-sdk main" | tee -a /etc/apt/sources.list.d/google-cloud-sdk.list && curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key --keyring /usr/share/keyrings/cloud.google.gpg  add - && apt-get update -y && apt-get install google-cloud-cli && apt-get install google-cloud-sdk-gke-gcloud-auth-plugin -y

RUN brew install ko
RUN wget -q https://raw.githubusercontent.com/dapr/cli/master/install/install.sh -O - | /bin/bash