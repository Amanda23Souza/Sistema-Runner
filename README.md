# Sistema Runner

Especificação do trabalho prático da disciplina Implementação e Integração (2026-01). Esta é a orientação para o que precisa ser feito:

- [Especificação](especificacao.md)
- [Design](design.md)

# Um plano preliminar de implementação...

Assumindo que os requisitos estão compreendidos (suposição provavelmente falsa) e que o projeto fornecido são suficientes para orientar a implementação, temos uma especulação inicial ([brainstorming](./docs/implementacao_transcricao.md)). Uma revisão segue abaixo com algumas 
especulações feitas em sala de aula:

- Para implementar os CLIs uma opção que oferece _cross-compiling_ nativo é a linguagem Go.
- A lista fornecida é um esforço de compreensão inicial do que deve ser feito. A ordem provavelmente será outra. Não há nenhum compromisso com a ordem (por enquanto).
  
Na ordem trabalhada:

1. Compreender o contexto do que deve ser produzido.
   1. Temos 2 CLIs e uma aplicação Java a ser construída. 
   2. A aplicação Java interage com dispositivos criptográfico (PKCS#11). Isso exige integração com biblioteca que faz a ponte entre Java e o dispositivo. Ou seja, usar SunPKCS11 (bridge) como "ponte" entre a aplicação Java e a biblioteca nativa do dispositivo. Observe que o dispositivo em si não está disponível, mas sim a biblioteca nativa dele (driver do fabricante) a partir da qual, de fato, o dispositivo é acessado. 
   3. A outra função da aplicação é validar os parâmetros de entrada. Naturalmente, precisam ser investigados e definidos. A definição deve ser sucedida por _design_ de interação com o usuário, apesar de CLI, é preciso projetar as opções. Veja tarefa abaixo incluída para esta finalidade.
   4. Para a simulação, assuma a existência da interface `SignatureService` com os métodos:
      - `sign(String message, String privateKey)`
      - `validate(String message, String signature, String publicKey)`
      - Parâmetros ainda são desconhecidos. Ou seja, acima tem apenas um "rumo". Por exemplo, quando se usa dispositivo criptográfico, não há "chave privada" acessível, conforme é suposto na assinatura deste método.
   5. Na figura (foto tirada), há um MC de Material Criptográfico, por exemplo, usb token ou smartcard, noutras palavras, o dispositivo físico onde se encontra a chave privada de certificado ICP-Brasil do tipo A3. Em tempo, esta chave "nunca" sai do dispositivo. Cada fabricante possui um "driver" de acesso ao MC, acessível de forma padronizada por uma biblioteca Java amplamente empregada, a SunPKCS11. Desta forma, esta biblioteca que acompanha o JDK é como a API JDBC, enquanto cada driver de um fornecedor seria o driver do Oracle, MySQL e assim por diante. Ou seja, uma estratégia recorrente.
2. Entradas e saídas deverão ser investigadas. Quais são os itens de dados de entrada? Quais os de saída? Como fornecê-los adequadamente? Lista de parâmetros (linha de comandos com flags, arquivos JSON ou outro, ou combinação)?
3. Protótipo Go (tenho insegurança sobre como fazer o que precisa ser feito pelo CLI em Go).
   1. Como lidar com parâmetros (cli)?
   2. Como iniciar processos em Go? (a aplicação em Java precisa ser iniciada e acompanhada)
   3. Como efetuar requisições via http (versão server do assinador)?
   4. Provavelmente todas estas operações podem ser produzidas por Modelo de IA corrente sem tanto esforço, contudo, é preciso um projeto adequado que admita testes.
4. Simulador. A implementação da interface `SignatureService` é o próprio processo de simulação. Sugestão de classe fake para implementar esta interface: `FakeSignatureService`.  
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
     

# O que está rolando... (desde 18/03/2026)

- Agora é o momento do planejamento da construção, ou pelo menos da primeira iteração. O que você vai fazer?
- Veja [aqui](./docs/planejamento.md) alguma orientação.

# O que está rolando... (desde 11/03/2026)

- O Princípio de Kerckhoffs diz que: "um sistema criptográfico deve permanecer seguro mesmo que tudo sobre o sistema seja público, exceto a chave privada".

# O que está rolando... (desde 10/03/2026)

- No primeiro encontra a [especificação](https://github.com/kyriosdata/runner/blob/v0.0.1/contexto.md) continha, por exemplo, requisitos sendo tratados como objetivos específicos, logo no início. Isso tinha que mudar. Na versão [melhorada](https://github.com/kyriosdata/runner/blob/v0.0.2/contexto.md), as seções foram alteradas e requisitos foram definidos na forma de user stories.

- Contudo, tenho 100% de certeza que ainda há muito para melhorar, inclusive na compreensão do próprio problema, antes mesmo até de trabalhar com uma estratégia como [SMART](https://thebaguide.com/blog/a-good-requirement-is-a-smart-requirement/) ou [INVEST](https://www.boost.co.nz/blog/2021/10/invest-criteria) para ajudar na caracterização dos requisitos. 

- Na versão v0.0.2 vemos critérios de aceitação, o que está alinhado com o BDD (Behavior Driven Development). Você pode consultar BDD na perspectiva de uma ferramenta concreta e real, o [Cucumber](https://cucumber.io/docs/).

- Apesar dos critérios, ainda não há uma definição clara de "done" para cada requisito, o que é fundamental. Esta definição de "done" é chamada, muitas vezes, de DoD (Definition of Done). Não ter ainda esta definição é natural, pois os requisitos ainda não atendem ao DoR (Definition of Ready), ou seja, ainda não estão prontos, conforme já mencionado.

- Quando olhamos para o [documento](https://github.com/kyyriosdata/runner/blob/v0.0.2/contexto.md), vemos que ele reúne requisitos e design. Em consequência, vamos dividir isso em dois documentos na v0.0.3. 

- Em tempo, conforme o SWEBOK, o que é considerado construção depende do modelo de ciclo de vida adotado, por exemplo, em modelos mais lineares, construção é precedida por requisitos e design, e sucedida por testes. Embora em muitos casos inclua codificação e depuração, também envolve planejamento, projeto detalhado, testes de unidade e testes de integração. 
