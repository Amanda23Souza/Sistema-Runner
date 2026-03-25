# Plano de implementação (revisado em 24/03/2026)

- CLIs serão feitos em Go por nativamente lidar com cross-compiling e funcionalidades usadas são "básicas" (Go 1.25).
- O assinador.jar será, por restrição de projeto, feito em Java 21.

Elementos de maior risco (desconhecidos):  

- Integração com Material Criptográfico via Java usando SunPKCS11 para esta finalidade. Talvez exista um simulador útil para testes.  Para simulação um indicação é [SoftHSM2](https://github.com/softhsm/SoftHSMv2).

- Compreender o contexto do que deve ser produzido.
- Identificar os parâmetros necessários tanto para criação quanto para validação de assinatura digital.
- Realizar o _design_ dos parâmetros. Como fornecê-los? Quais os flags? Arquivo?

- Definir a interface `SignatureService` com os métodos:
      - `sign(a serem definidos)`
      - `validate(a serem definidos)`
- A implementação desta interface `SignatureService` é a implementação da simulação a ser realizada pela classe `FakeSignatureService`.

- Protótipo Go
   1. Como lidar com parâmetros (cli)? Usar [Cobra](https://cobra.dev/)?
   2. Como iniciar processos em Go? (a aplicação em Java precisa ser iniciada e acompanhada)
   3. Como efetuar requisições via http?

4. Simulador.   
5. A interface da foto é substituída aqui por `SignatureService` conforme acima. 
6. O modo server é melhor descrito como uma aplicação web, que oferece endpoints para assinatura e validação de documentos. Ou seja, é necessário um controller `SignatureController` com a definição dos endpoints. Na foto são definidos `/sign` e `/validate`.
7. Suponho que seja criar o projeto Java (esqueleto, pom.xml, ...)
8. A qualidade da anotação é relevante, simplesmente não me lembro nem o 7 nem o 8, apesar de, em algum momento, estar "óbvio". 
9. Ao iniciar a aplicação Java é oportuno indicar qual porta usar, isso para evitar conflito com portas já em uso, ou seja, detectar se a porta padrão está em uso por outra aplicação e, caso esteja, identificar outra disponível. Assumi que isso pode ser feito "tranquilamente" em Go. Isso é uma suposição.
10. É preciso para o processo iniciado via Go. Talvez possa ser incorporado ao protótipo citado anteriormente. Ou seja, há um conjunto de operações a serem realizadas pelo CLI em Go para gerenciar a aplicação Java.
11. Banco de dados. Precisamos de um banco de dados para armazenar os dados necessários para as operações do Sistema Runner, por exemplo, o runtime Java empregado pelo CLI, a porta empregada pelo processo em execução, o PID do processo em execução e outras. Lembre-se de que o CLI deve baixar o runtime java, o que significa usar o sistema de arquivos local, desempacotá-lo e disponibilizá-lo para uso local. Isso pressupõe o uso de um diretório, por exemplo, `.hubsaude` na home dir do usuário em questão, dentro deste diretório onde depositar o runtime java descompactado, o arquivo contendo informações sobre processos, e outras.
12. É preciso realizar o download da aplicação Java no caso do simulador.jar. Embora o CLI deva ter essa url armazenada internamente (hardwired), a opção `--source` ou equivalente deve ser possível para permitir que a busca possa ocorrer em outro local sem necessidade de atualização do CLI. 
13. Startup. Processo de inicialização dos CLI. Cada um tem suas especificidades, mas ambos, por exemplo, dependem do runtime java e do banco de dados. Então é preciso ir buscar informações no banco de dados, que pode não estar disponível então teria que ser criado (por exemplo, esse é o cenário inicial), mas pode estar disponível mas sem o runtime, então teria que baixar o runtime, descompactar e assim por diante. É uma sequencia de passos que deve ser bem projetada, pois isso é relevante para a percepção do usuário (desempenho).
14. Scripts para CI/CD. Ou seja, configura o GitHub Actions para integração contínua (CI), compilação e execução de testes de unidade e integração, seguida da entrega contínua que, neste caso, limita-se a disponibilizar as aplicações geradas (os clis nas várias plataformas e a aplicação Java criada, o assinador).
