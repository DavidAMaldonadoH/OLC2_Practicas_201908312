
//----------------------------------------------------
// The following code was generated by CUP v0.11b 20160615 (GIT 4ac7450)
//----------------------------------------------------

package AnalizadorBinario;

import java_cup.runtime.*;
import java_cup.runtime.XMLElement;

/** CUP v0.11b 20160615 (GIT 4ac7450) generated parser.
  */
@SuppressWarnings({"rawtypes"})
public class SintacticoBinario extends java_cup.runtime.lr_parser {

 public final Class getSymbolContainer() {
    return sym.class;
}

  /** Default constructor. */
  @Deprecated
  public SintacticoBinario() {super();}

  /** Constructor which sets the default scanner. */
  @Deprecated
  public SintacticoBinario(java_cup.runtime.Scanner s) {super(s);}

  /** Constructor which sets the default scanner. */
  public SintacticoBinario(java_cup.runtime.Scanner s, java_cup.runtime.SymbolFactory sf) {super(s,sf);}

  /** Production table. */
  protected static final short _production_table[][] = 
    unpackFromStrings(new String[] {
    "\000\006\000\002\002\004\000\002\002\003\000\002\003" +
    "\004\000\002\003\004\000\002\003\003\000\002\003\003" +
    "" });

  /** Access to production table. */
  public short[][] production_table() {return _production_table;}

  /** Parse-action table. */
  protected static final short[][] _action_table = 
    unpackFromStrings(new String[] {
    "\000\010\000\006\004\004\005\005\001\002\000\010\002" +
    "\ufffd\004\004\005\005\001\002\000\010\002\ufffc\004\004" +
    "\005\005\001\002\000\004\002\010\001\002\000\004\002" +
    "\000\001\002\000\004\002\001\001\002\000\004\002\ufffe" +
    "\001\002\000\004\002\uffff\001\002" });

  /** Access to parse-action table. */
  public short[][] action_table() {return _action_table;}

  /** <code>reduce_goto</code> table. */
  protected static final short[][] _reduce_table = 
    unpackFromStrings(new String[] {
    "\000\010\000\006\002\005\003\006\001\001\000\004\003" +
    "\011\001\001\000\004\003\010\001\001\000\002\001\001" +
    "\000\002\001\001\000\002\001\001\000\002\001\001\000" +
    "\002\001\001" });

  /** Access to <code>reduce_goto</code> table. */
  public short[][] reduce_table() {return _reduce_table;}

  /** Instance of action encapsulation class. */
  protected CUP$SintacticoBinario$actions action_obj;

  /** Action encapsulation object initializer. */
  protected void init_actions()
    {
      action_obj = new CUP$SintacticoBinario$actions(this);
    }

  /** Invoke a user supplied parse action. */
  public java_cup.runtime.Symbol do_action(
    int                        act_num,
    java_cup.runtime.lr_parser parser,
    java.util.Stack            stack,
    int                        top)
    throws java.lang.Exception
  {
    /* call code in generated class */
    return action_obj.CUP$SintacticoBinario$do_action(act_num, parser, stack, top);
  }

  /** Indicates start state. */
  public int start_state() {return 0;}
  /** Indicates start production. */
  public int start_production() {return 0;}

  /** <code>EOF</code> Symbol index. */
  public int EOF_sym() {return 0;}

  /** <code>error</code> Symbol index. */
  public int error_sym() {return 1;}



    int contador0 = 0;
    int contador1 = 0;
    
    public void syntax_error(Symbol s){ 
        System.out.println("Error Sint??ctico en la L??nea " + (s.left) +
        " Columna "+s.right+ ". No se esperaba este componente: " +s.value+"."); 
    } 

    public void unrecovered_syntax_error(Symbol s) throws java.lang.Exception{ 
        System.out.println("Error s??ntactico irrecuperable en la L??nea " + 
        (s.left)+ " Columna "+s.right+". Componente " + s.value + 
        " no reconocido."); 
    }  



/** Cup generated class to encapsulate user supplied action code.*/
@SuppressWarnings({"rawtypes", "unchecked", "unused"})
class CUP$SintacticoBinario$actions {
  private final SintacticoBinario parser;

  /** Constructor */
  CUP$SintacticoBinario$actions(SintacticoBinario parser) {
    this.parser = parser;
  }

  /** Method 0 with the actual generated action code for actions 0 to 300. */
  public final java_cup.runtime.Symbol CUP$SintacticoBinario$do_action_part00000000(
    int                        CUP$SintacticoBinario$act_num,
    java_cup.runtime.lr_parser CUP$SintacticoBinario$parser,
    java.util.Stack            CUP$SintacticoBinario$stack,
    int                        CUP$SintacticoBinario$top)
    throws java.lang.Exception
    {
      /* Symbol object for return from actions */
      java_cup.runtime.Symbol CUP$SintacticoBinario$result;

      /* select the action based on the action number */
      switch (CUP$SintacticoBinario$act_num)
        {
          /*. . . . . . . . . . . . . . . . . . . .*/
          case 0: // $START ::= S EOF 
            {
              Object RESULT =null;
		int start_valleft = ((java_cup.runtime.Symbol)CUP$SintacticoBinario$stack.elementAt(CUP$SintacticoBinario$top-1)).left;
		int start_valright = ((java_cup.runtime.Symbol)CUP$SintacticoBinario$stack.elementAt(CUP$SintacticoBinario$top-1)).right;
		Object start_val = (Object)((java_cup.runtime.Symbol) CUP$SintacticoBinario$stack.elementAt(CUP$SintacticoBinario$top-1)).value;
		RESULT = start_val;
              CUP$SintacticoBinario$result = parser.getSymbolFactory().newSymbol("$START",0, ((java_cup.runtime.Symbol)CUP$SintacticoBinario$stack.elementAt(CUP$SintacticoBinario$top-1)), ((java_cup.runtime.Symbol)CUP$SintacticoBinario$stack.peek()), RESULT);
            }
          /* ACCEPT */
          CUP$SintacticoBinario$parser.done_parsing();
          return CUP$SintacticoBinario$result;

          /*. . . . . . . . . . . . . . . . . . . .*/
          case 1: // S ::= N 
            {
              Object RESULT =null;
		 
    System.out.println("> Salida: " + contador1 + " veces 1, " + contador0 + " veces 0");

              CUP$SintacticoBinario$result = parser.getSymbolFactory().newSymbol("S",0, ((java_cup.runtime.Symbol)CUP$SintacticoBinario$stack.peek()), ((java_cup.runtime.Symbol)CUP$SintacticoBinario$stack.peek()), RESULT);
            }
          return CUP$SintacticoBinario$result;

          /*. . . . . . . . . . . . . . . . . . . .*/
          case 2: // N ::= UNO N 
            {
              Object RESULT =null;
		 contador1 += 1; 
              CUP$SintacticoBinario$result = parser.getSymbolFactory().newSymbol("N",1, ((java_cup.runtime.Symbol)CUP$SintacticoBinario$stack.elementAt(CUP$SintacticoBinario$top-1)), ((java_cup.runtime.Symbol)CUP$SintacticoBinario$stack.peek()), RESULT);
            }
          return CUP$SintacticoBinario$result;

          /*. . . . . . . . . . . . . . . . . . . .*/
          case 3: // N ::= CERO N 
            {
              Object RESULT =null;
		 contador0 += 1; 
              CUP$SintacticoBinario$result = parser.getSymbolFactory().newSymbol("N",1, ((java_cup.runtime.Symbol)CUP$SintacticoBinario$stack.elementAt(CUP$SintacticoBinario$top-1)), ((java_cup.runtime.Symbol)CUP$SintacticoBinario$stack.peek()), RESULT);
            }
          return CUP$SintacticoBinario$result;

          /*. . . . . . . . . . . . . . . . . . . .*/
          case 4: // N ::= UNO 
            {
              Object RESULT =null;
		 contador1 = 1; 
              CUP$SintacticoBinario$result = parser.getSymbolFactory().newSymbol("N",1, ((java_cup.runtime.Symbol)CUP$SintacticoBinario$stack.peek()), ((java_cup.runtime.Symbol)CUP$SintacticoBinario$stack.peek()), RESULT);
            }
          return CUP$SintacticoBinario$result;

          /*. . . . . . . . . . . . . . . . . . . .*/
          case 5: // N ::= CERO 
            {
              Object RESULT =null;
		 contador0 = 1; 
              CUP$SintacticoBinario$result = parser.getSymbolFactory().newSymbol("N",1, ((java_cup.runtime.Symbol)CUP$SintacticoBinario$stack.peek()), ((java_cup.runtime.Symbol)CUP$SintacticoBinario$stack.peek()), RESULT);
            }
          return CUP$SintacticoBinario$result;

          /* . . . . . .*/
          default:
            throw new Exception(
               "Invalid action number "+CUP$SintacticoBinario$act_num+"found in internal parse table");

        }
    } /* end of method */

  /** Method splitting the generated action code into several parts. */
  public final java_cup.runtime.Symbol CUP$SintacticoBinario$do_action(
    int                        CUP$SintacticoBinario$act_num,
    java_cup.runtime.lr_parser CUP$SintacticoBinario$parser,
    java.util.Stack            CUP$SintacticoBinario$stack,
    int                        CUP$SintacticoBinario$top)
    throws java.lang.Exception
    {
              return CUP$SintacticoBinario$do_action_part00000000(
                               CUP$SintacticoBinario$act_num,
                               CUP$SintacticoBinario$parser,
                               CUP$SintacticoBinario$stack,
                               CUP$SintacticoBinario$top);
    }
}

}
