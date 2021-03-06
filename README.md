# <h1 align ='center'>---------- ProjetosGo ---------- </h1>
Meus repositórios de projetos projetos da linguagem de programação Go.
Estes repositórios são apenas testes e algumas anotações durante os meus estudos dessa linguagem
#
<h2 align='center'> ----------- Projetos ---------- </h2>

<!------------VARIÁVEIS------------>
<h3 align='center'> Variaveis </h3>
Neste repositório existem um arquivo .go tratando sobre os assuntos:
<br> * Formas de declarar variáveis
<br> * Blocos de Variáveis
<br> * Converção de variáveis

<!------------DATATYPES------------>
<h3 align='center'> DataTypes </h3>
Neste repositório existem um arquivo .go tratando sobre os assuntos: <br>
<!------Boolean------>
<br> <h4 align='center'> * Boolean </h4>
  <ul>
    <li> Valores são True (verdadeiro) ou False (falso) </li>
    <li> Valor zerado é falso (false)             </li>
  </ul>
<!------Tipos Numéricos-------->
<br> <h4 align='center'> * Tipos Numéricos </h4>

  <!--Signed Integers-->
  <ul> <h5> Signed Integers </h4>
    <li> Int tem o tamanho variado, mas o minimo é 32 bits  </li>
    <li> 8bit(int8) até 64bit (int64)                       </li>
  </ul>
  
  <!--Unsigned Integers-->
  <ul> <h5> Unsigned Integers </h4>
    <li> 8 bit (byte e uint8) até 32 bit (uint32) </li>
  </ul>
  <ul> <h5> Operações aritmétricas dos números integers </h4>
    <li> Adição         </li>
    <li> Subtração      </li>
    <li> Multiplicação  </li>
    <li> Divisão        </li>
    <li> Restante       </li>
  </ul>
  
  <!--Bitwise-->
  <ul> <h5> Operadores bitwise </h4>
    <li> And  </li>
    <li> Or   </li>
    <li> Not  </li>
  </ul>
  
  <!--Float-->
  <ul> <h5> Float </h4>
    <li> Seguem o padrão IEEE-754 </li>
    <li> Valor zerado é 0         </li>
    <li> Versões 32 e 64 bits     </li>
  </ul>
    <ul> <h5> Operações aritmétricas dos números floats </h4>
    <li> Adição         </li>
    <li> Subtração      </li>
    <li> Multiplicação  </li>
    <li> Divisão        </li>
  </ul>
  
  <!--Complex-->
  <ul> <h5> Complex </h4>
    <li> Valor zerado é (0 + 0i)  </li>
    <li> Versões 64 e 128 bits    </li>
  </ul>
 
  <!--Tipos de texto-->
  <br> <h4 align='center'> * Tipos de Texto </h4>
    <ul> <h5> Strings </h4>
      <li> UTF-8                                            </li>
      <li> Imutável                                         </li>
      <li> Pode ser contatenado com o operador de soma (+)  </li>
      <li> Pode ser convertido para binario                 </li>
    </ul>
    <ul> <h5> Rune </h4>
      <li> UTF-32 </li>
      <li> int32  </li>
    </ul>
    
<!------------Constantes------------>
<h3 align='center'> Constantes </h3>

<!----------Arrays---------->
<h3 align='center'> Arrays </h3>
  <ul>
    <li> Coleção de itens com o mesmo tipo        </li>
    <li> Tamanho fixo                             </li>
    <li> A função len retorna o tamanho do array  </li>
  </ul>
  
  <!----------Slices---------->
  <h3 align='center'> Slices </h3>
  <ul>
    <li> Suportado por arrays                         </li>
    <li> A função len retorna o tamanho do slice      </li>
    <li> A função cap retorna a capacidade do slice   </li>
    <li> A função append adiciona itens para o slice  </li>
  </ul>
  <!----------Maps---------->
  <h3 align='center'> Maps </h3>
  <ul>
    <li> Coleção de dados que podem ser acessados via "Keys" </li>
    <li> Criados via literais ou pela função Make </li>
    <li> Checa a presença com a forma de resultado "value, ok" </li>
    <li> Multiplas atribuições referecem se ao mesmo map </li>
  </ul>
  
  <!----------Structs---------->
  <h3 align='center'> Structs </h3>
  <ul>
    <li> Coleção de dados que descrevem o mesmo conceito </li>
    <li> Organizados por campos de nomes </li>
    <li>Normalmente criados como tipos mas structs anonimas podem ser criadas </li>
    <li> São tipos de valores de dados </li>
    <li> Não existe herança, mas pode ser usada composição via embedding </li>
  </ul>
