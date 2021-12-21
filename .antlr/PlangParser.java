// Generated from /home/mo/opnlab/v3/Plang.g4 by ANTLR 4.8
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class PlangParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, T__15=16, T__16=17, 
		T__17=18, T__18=19, T__19=20, T__20=21, T__21=22, T__22=23, T__23=24, 
		T__24=25, T__25=26, T__26=27, T__27=28, T__28=29, T__29=30, T__30=31, 
		T__31=32, T__32=33, T__33=34, T__34=35, T__35=36, T__36=37, T__37=38, 
		T__38=39, Ident=40, Int=41, Str=42, Float=43, WS=44;
	public static final int
		RULE_program = 0, RULE_block = 1, RULE_ident = 2, RULE_intType = 3, RULE_floatType = 4, 
		RULE_stringType = 5, RULE_mathSign = 6, RULE_nestable = 7, RULE_stmt = 8, 
		RULE_fnArg = 9, RULE_fnCall = 10, RULE_objectField = 11, RULE_expr = 12;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "block", "ident", "intType", "floatType", "stringType", "mathSign", 
			"nestable", "stmt", "fnArg", "fnCall", "objectField", "expr"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'+'", "'-'", "'/'", "'*'", "'.'", "'['", "']'", "'import'", "'as'", 
			"'pub'", "':='", "'='", "'func'", "'('", "','", "'...'", "')'", "'{'", 
			"'}'", "'for'", "'in'", "'else'", "'if'", "'return'", "'break'", "';'", 
			"':'", "'true'", "'false'", "'null'", "'!'", "'=='", "'!='", "'>'", "'<='", 
			"'<'", "'>='", "'++'", "'--'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, "Ident", "Int", "Str", "Float", "WS"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "Plang.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public PlangParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class ProgramContext extends ParserRuleContext {
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode EOF() { return getToken(PlangParser.EOF, 0); }
		public ProgramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_program; }
	}

	public final ProgramContext program() throws RecognitionException {
		ProgramContext _localctx = new ProgramContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_program);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(26);
			block();
			setState(27);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class BlockContext extends ParserRuleContext {
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public BlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block; }
	}

	public final BlockContext block() throws RecognitionException {
		BlockContext _localctx = new BlockContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_block);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(32);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__1) | (1L << T__5) | (1L << T__7) | (1L << T__9) | (1L << T__12) | (1L << T__13) | (1L << T__17) | (1L << T__19) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << Ident) | (1L << Int) | (1L << Str) | (1L << Float))) != 0)) {
				{
				{
				setState(29);
				stmt();
				}
				}
				setState(34);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class IdentContext extends ParserRuleContext {
		public TerminalNode Ident() { return getToken(PlangParser.Ident, 0); }
		public IdentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ident; }
	}

	public final IdentContext ident() throws RecognitionException {
		IdentContext _localctx = new IdentContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_ident);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(35);
			match(Ident);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class IntTypeContext extends ParserRuleContext {
		public TerminalNode Int() { return getToken(PlangParser.Int, 0); }
		public IntTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_intType; }
	}

	public final IntTypeContext intType() throws RecognitionException {
		IntTypeContext _localctx = new IntTypeContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_intType);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(37);
			match(Int);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FloatTypeContext extends ParserRuleContext {
		public TerminalNode Float() { return getToken(PlangParser.Float, 0); }
		public FloatTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_floatType; }
	}

	public final FloatTypeContext floatType() throws RecognitionException {
		FloatTypeContext _localctx = new FloatTypeContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_floatType);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(39);
			match(Float);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class StringTypeContext extends ParserRuleContext {
		public TerminalNode Str() { return getToken(PlangParser.Str, 0); }
		public StringTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_stringType; }
	}

	public final StringTypeContext stringType() throws RecognitionException {
		StringTypeContext _localctx = new StringTypeContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_stringType);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(41);
			match(Str);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class MathSignContext extends ParserRuleContext {
		public MathSignContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_mathSign; }
	}

	public final MathSignContext mathSign() throws RecognitionException {
		MathSignContext _localctx = new MathSignContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_mathSign);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(43);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << T__1) | (1L << T__2) | (1L << T__3))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class NestableContext extends ParserRuleContext {
		public IdentContext DOT;
		public ExprContext INDEX;
		public IdentContext ident() {
			return getRuleContext(IdentContext.class,0);
		}
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public NestableContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_nestable; }
	}

	public final NestableContext nestable() throws RecognitionException {
		NestableContext _localctx = new NestableContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_nestable);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(51);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__4:
				{
				{
				setState(45);
				match(T__4);
				setState(46);
				((NestableContext)_localctx).DOT = ident();
				}
				}
				break;
			case T__5:
				{
				{
				setState(47);
				match(T__5);
				setState(48);
				((NestableContext)_localctx).INDEX = expr(0);
				setState(49);
				match(T__6);
				}
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class StmtContext extends ParserRuleContext {
		public StmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_stmt; }
	 
		public StmtContext() { }
		public void copyFrom(StmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class ForInStmtContext extends StmtContext {
		public IdentContext ident;
		public List<IdentContext> NAMES = new ArrayList<IdentContext>();
		public ExprContext TARGET;
		public BlockContext BODY;
		public List<IdentContext> ident() {
			return getRuleContexts(IdentContext.class);
		}
		public IdentContext ident(int i) {
			return getRuleContext(IdentContext.class,i);
		}
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public ForInStmtContext(StmtContext ctx) { copyFrom(ctx); }
	}
	public static class FnCallStmtContext extends StmtContext {
		public FnCallContext fnCall() {
			return getRuleContext(FnCallContext.class,0);
		}
		public FnCallStmtContext(StmtContext ctx) { copyFrom(ctx); }
	}
	public static class ImportStmtContext extends StmtContext {
		public StringTypeContext NAME;
		public IdentContext ALIAS;
		public StringTypeContext stringType() {
			return getRuleContext(StringTypeContext.class,0);
		}
		public IdentContext ident() {
			return getRuleContext(IdentContext.class,0);
		}
		public ImportStmtContext(StmtContext ctx) { copyFrom(ctx); }
	}
	public static class IfStmtContext extends StmtContext {
		public ExprContext COND;
		public BlockContext BODY;
		public ExprContext expr;
		public List<ExprContext> ELIFCOND = new ArrayList<ExprContext>();
		public BlockContext block;
		public List<BlockContext> ELIFBODY = new ArrayList<BlockContext>();
		public BlockContext ELSEBODY;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public List<BlockContext> block() {
			return getRuleContexts(BlockContext.class);
		}
		public BlockContext block(int i) {
			return getRuleContext(BlockContext.class,i);
		}
		public IfStmtContext(StmtContext ctx) { copyFrom(ctx); }
	}
	public static class ExprStmtContext extends StmtContext {
		public ExprContext E;
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public ExprStmtContext(StmtContext ctx) { copyFrom(ctx); }
	}
	public static class AssignStmtContext extends StmtContext {
		public IdentContext ROOT;
		public NestableContext nestable;
		public List<NestableContext> CHAIN = new ArrayList<NestableContext>();
		public ExprContext VALUE;
		public IdentContext ident() {
			return getRuleContext(IdentContext.class,0);
		}
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public List<NestableContext> nestable() {
			return getRuleContexts(NestableContext.class);
		}
		public NestableContext nestable(int i) {
			return getRuleContext(NestableContext.class,i);
		}
		public AssignStmtContext(StmtContext ctx) { copyFrom(ctx); }
	}
	public static class BreakStmtContext extends StmtContext {
		public BreakStmtContext(StmtContext ctx) { copyFrom(ctx); }
	}
	public static class DeclareStmtContext extends StmtContext {
		public Token PUB;
		public IdentContext NAME;
		public ExprContext VALUE;
		public IdentContext ident() {
			return getRuleContext(IdentContext.class,0);
		}
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public DeclareStmtContext(StmtContext ctx) { copyFrom(ctx); }
	}
	public static class FnDeclareStmtContext extends StmtContext {
		public IdentContext NAME;
		public IdentContext ident;
		public List<IdentContext> ARGS = new ArrayList<IdentContext>();
		public IdentContext RESTARG;
		public BlockContext BODY;
		public List<IdentContext> ident() {
			return getRuleContexts(IdentContext.class);
		}
		public IdentContext ident(int i) {
			return getRuleContext(IdentContext.class,i);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public FnDeclareStmtContext(StmtContext ctx) { copyFrom(ctx); }
	}
	public static class ForCondStmtContext extends StmtContext {
		public ExprContext COND;
		public BlockContext BODY;
		public BlockContext ELSEBODY;
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public List<BlockContext> block() {
			return getRuleContexts(BlockContext.class);
		}
		public BlockContext block(int i) {
			return getRuleContext(BlockContext.class,i);
		}
		public ForCondStmtContext(StmtContext ctx) { copyFrom(ctx); }
	}
	public static class ReturnStmtContext extends StmtContext {
		public ExprContext VALUE;
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public ReturnStmtContext(StmtContext ctx) { copyFrom(ctx); }
	}

	public final StmtContext stmt() throws RecognitionException {
		StmtContext _localctx = new StmtContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_stmt);
		int _la;
		try {
			int _alt;
			setState(161);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,14,_ctx) ) {
			case 1:
				_localctx = new ImportStmtContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(53);
				match(T__7);
				setState(54);
				((ImportStmtContext)_localctx).NAME = stringType();
				setState(57);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__8) {
					{
					setState(55);
					match(T__8);
					setState(56);
					((ImportStmtContext)_localctx).ALIAS = ident();
					}
				}

				}
				break;
			case 2:
				_localctx = new DeclareStmtContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(60);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__9) {
					{
					setState(59);
					((DeclareStmtContext)_localctx).PUB = match(T__9);
					}
				}

				setState(62);
				((DeclareStmtContext)_localctx).NAME = ident();
				setState(63);
				match(T__10);
				setState(64);
				((DeclareStmtContext)_localctx).VALUE = expr(0);
				}
				break;
			case 3:
				_localctx = new AssignStmtContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(66);
				((AssignStmtContext)_localctx).ROOT = ident();
				setState(70);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__4 || _la==T__5) {
					{
					{
					setState(67);
					((AssignStmtContext)_localctx).nestable = nestable();
					((AssignStmtContext)_localctx).CHAIN.add(((AssignStmtContext)_localctx).nestable);
					}
					}
					setState(72);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(73);
				match(T__11);
				setState(74);
				((AssignStmtContext)_localctx).VALUE = expr(0);
				}
				break;
			case 4:
				_localctx = new FnCallStmtContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(76);
				fnCall();
				}
				break;
			case 5:
				_localctx = new FnDeclareStmtContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(77);
				match(T__12);
				setState(78);
				((FnDeclareStmtContext)_localctx).NAME = ident();
				setState(79);
				match(T__13);
				setState(97);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__15 || _la==Ident) {
					{
					setState(95);
					_errHandler.sync(this);
					switch (_input.LA(1)) {
					case Ident:
						{
						{
						setState(80);
						((FnDeclareStmtContext)_localctx).ident = ident();
						((FnDeclareStmtContext)_localctx).ARGS.add(((FnDeclareStmtContext)_localctx).ident);
						setState(85);
						_errHandler.sync(this);
						_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
						while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
							if ( _alt==1 ) {
								{
								{
								setState(81);
								match(T__14);
								setState(82);
								((FnDeclareStmtContext)_localctx).ident = ident();
								((FnDeclareStmtContext)_localctx).ARGS.add(((FnDeclareStmtContext)_localctx).ident);
								}
								} 
							}
							setState(87);
							_errHandler.sync(this);
							_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
						}
						setState(91);
						_errHandler.sync(this);
						_la = _input.LA(1);
						if (_la==T__14) {
							{
							setState(88);
							match(T__14);
							setState(89);
							match(T__15);
							setState(90);
							((FnDeclareStmtContext)_localctx).RESTARG = ident();
							}
						}

						}
						}
						break;
					case T__15:
						{
						{
						setState(93);
						match(T__15);
						setState(94);
						((FnDeclareStmtContext)_localctx).RESTARG = ident();
						}
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					}
				}

				setState(99);
				match(T__16);
				setState(100);
				match(T__17);
				setState(101);
				((FnDeclareStmtContext)_localctx).BODY = block();
				setState(102);
				match(T__18);
				}
				break;
			case 6:
				_localctx = new ForInStmtContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(104);
				match(T__19);
				setState(105);
				((ForInStmtContext)_localctx).ident = ident();
				((ForInStmtContext)_localctx).NAMES.add(((ForInStmtContext)_localctx).ident);
				setState(110);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__14) {
					{
					{
					setState(106);
					match(T__14);
					setState(107);
					((ForInStmtContext)_localctx).ident = ident();
					((ForInStmtContext)_localctx).NAMES.add(((ForInStmtContext)_localctx).ident);
					}
					}
					setState(112);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(113);
				match(T__20);
				setState(114);
				((ForInStmtContext)_localctx).TARGET = expr(0);
				setState(115);
				match(T__17);
				setState(116);
				((ForInStmtContext)_localctx).BODY = block();
				setState(117);
				match(T__18);
				}
				break;
			case 7:
				_localctx = new ForCondStmtContext(_localctx);
				enterOuterAlt(_localctx, 7);
				{
				setState(119);
				match(T__19);
				setState(120);
				((ForCondStmtContext)_localctx).COND = expr(0);
				setState(121);
				match(T__17);
				setState(122);
				((ForCondStmtContext)_localctx).BODY = block();
				setState(123);
				match(T__18);
				setState(129);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__21) {
					{
					setState(124);
					match(T__21);
					setState(125);
					match(T__17);
					setState(126);
					((ForCondStmtContext)_localctx).ELSEBODY = block();
					setState(127);
					match(T__18);
					}
				}

				}
				break;
			case 8:
				_localctx = new IfStmtContext(_localctx);
				enterOuterAlt(_localctx, 8);
				{
				setState(131);
				match(T__22);
				setState(132);
				((IfStmtContext)_localctx).COND = expr(0);
				setState(133);
				match(T__17);
				setState(134);
				((IfStmtContext)_localctx).BODY = block();
				setState(135);
				match(T__18);
				setState(145);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,11,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(136);
						match(T__21);
						setState(137);
						match(T__22);
						setState(138);
						((IfStmtContext)_localctx).expr = expr(0);
						((IfStmtContext)_localctx).ELIFCOND.add(((IfStmtContext)_localctx).expr);
						setState(139);
						match(T__17);
						setState(140);
						((IfStmtContext)_localctx).block = block();
						((IfStmtContext)_localctx).ELIFBODY.add(((IfStmtContext)_localctx).block);
						setState(141);
						match(T__18);
						}
						} 
					}
					setState(147);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,11,_ctx);
				}
				setState(153);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__21) {
					{
					setState(148);
					match(T__21);
					setState(149);
					match(T__17);
					setState(150);
					((IfStmtContext)_localctx).ELSEBODY = block();
					setState(151);
					match(T__18);
					}
				}

				}
				break;
			case 9:
				_localctx = new ReturnStmtContext(_localctx);
				enterOuterAlt(_localctx, 9);
				{
				setState(155);
				match(T__23);
				setState(157);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
				case 1:
					{
					setState(156);
					((ReturnStmtContext)_localctx).VALUE = expr(0);
					}
					break;
				}
				}
				break;
			case 10:
				_localctx = new BreakStmtContext(_localctx);
				enterOuterAlt(_localctx, 10);
				{
				setState(159);
				match(T__24);
				}
				break;
			case 11:
				_localctx = new ExprStmtContext(_localctx);
				enterOuterAlt(_localctx, 11);
				{
				setState(160);
				((ExprStmtContext)_localctx).E = expr(0);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FnArgContext extends ParserRuleContext {
		public ExprContext ARG;
		public ExprContext SPREADABLE;
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public FnArgContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_fnArg; }
	}

	public final FnArgContext fnArg() throws RecognitionException {
		FnArgContext _localctx = new FnArgContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_fnArg);
		try {
			setState(167);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,15,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(163);
				((FnArgContext)_localctx).ARG = expr(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(164);
				((FnArgContext)_localctx).SPREADABLE = expr(0);
				setState(165);
				match(T__15);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FnCallContext extends ParserRuleContext {
		public IdentContext NAME;
		public ExprContext EXPR;
		public FnArgContext fnArg;
		public List<FnArgContext> ARGS = new ArrayList<FnArgContext>();
		public IdentContext ident() {
			return getRuleContext(IdentContext.class,0);
		}
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public List<FnArgContext> fnArg() {
			return getRuleContexts(FnArgContext.class);
		}
		public FnArgContext fnArg(int i) {
			return getRuleContext(FnArgContext.class,i);
		}
		public FnCallContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_fnCall; }
	}

	public final FnCallContext fnCall() throws RecognitionException {
		FnCallContext _localctx = new FnCallContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_fnCall);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(171);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,16,_ctx) ) {
			case 1:
				{
				setState(169);
				((FnCallContext)_localctx).NAME = ident();
				}
				break;
			case 2:
				{
				setState(170);
				((FnCallContext)_localctx).EXPR = expr(0);
				}
				break;
			}
			setState(173);
			match(T__13);
			setState(182);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__1) | (1L << T__5) | (1L << T__13) | (1L << T__17) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << Ident) | (1L << Int) | (1L << Str) | (1L << Float))) != 0)) {
				{
				setState(174);
				((FnCallContext)_localctx).fnArg = fnArg();
				((FnCallContext)_localctx).ARGS.add(((FnCallContext)_localctx).fnArg);
				setState(179);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__14) {
					{
					{
					setState(175);
					match(T__14);
					setState(176);
					((FnCallContext)_localctx).fnArg = fnArg();
					((FnCallContext)_localctx).ARGS.add(((FnCallContext)_localctx).fnArg);
					}
					}
					setState(181);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(184);
			match(T__16);
			setState(186);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__25) {
				{
				setState(185);
				match(T__25);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ObjectFieldContext extends ParserRuleContext {
		public IdentContext STATIC;
		public ExprContext DYNAMIC;
		public ExprContext VALUE;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public IdentContext ident() {
			return getRuleContext(IdentContext.class,0);
		}
		public ObjectFieldContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_objectField; }
	}

	public final ObjectFieldContext objectField() throws RecognitionException {
		ObjectFieldContext _localctx = new ObjectFieldContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_objectField);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(193);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Ident:
				{
				setState(188);
				((ObjectFieldContext)_localctx).STATIC = ident();
				}
				break;
			case T__5:
				{
				setState(189);
				match(T__5);
				setState(190);
				((ObjectFieldContext)_localctx).DYNAMIC = expr(0);
				setState(191);
				match(T__6);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(195);
			match(T__26);
			setState(196);
			((ObjectFieldContext)_localctx).VALUE = expr(0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExprContext extends ParserRuleContext {
		public ExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr; }
	 
		public ExprContext() { }
		public void copyFrom(ExprContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class BoolExprContext extends ExprContext {
		public BoolExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class StringExprContext extends ExprContext {
		public StringTypeContext stringType() {
			return getRuleContext(StringTypeContext.class,0);
		}
		public StringExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class EqualExprContext extends ExprContext {
		public ExprContext L;
		public ExprContext R;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public EqualExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class IdentExprContext extends ExprContext {
		public IdentContext ident() {
			return getRuleContext(IdentContext.class,0);
		}
		public IdentExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class DecExprContext extends ExprContext {
		public IdentContext ident() {
			return getRuleContext(IdentContext.class,0);
		}
		public DecExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class FloatExprContext extends ExprContext {
		public FloatTypeContext floatType() {
			return getRuleContext(FloatTypeContext.class,0);
		}
		public FloatExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class IncExprContext extends ExprContext {
		public ExprContext L;
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public IncExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class GtExprContext extends ExprContext {
		public ExprContext L;
		public ExprContext R;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public GtExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class LtExprContext extends ExprContext {
		public ExprContext L;
		public ExprContext R;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public LtExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class DotExprContext extends ExprContext {
		public ExprContext PARENT;
		public IdentContext CHILD;
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public IdentContext ident() {
			return getRuleContext(IdentContext.class,0);
		}
		public DotExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class NotEqualExprContext extends ExprContext {
		public ExprContext L;
		public ExprContext R;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public NotEqualExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class LteExprContext extends ExprContext {
		public ExprContext L;
		public ExprContext R;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public LteExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class IndexExprContext extends ExprContext {
		public ExprContext PARENT;
		public ExprContext CHILD;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public IndexExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class NullExprContext extends ExprContext {
		public NullExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class GteExprContext extends ExprContext {
		public ExprContext L;
		public ExprContext R;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public GteExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class GroupExprContext extends ExprContext {
		public ExprContext EXPR;
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public GroupExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class NegativeExprContext extends ExprContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public NegativeExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class ArrayExprContext extends ExprContext {
		public ExprContext expr;
		public List<ExprContext> ELMS = new ArrayList<ExprContext>();
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public ArrayExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class ObjectExprContext extends ExprContext {
		public ObjectFieldContext objectField;
		public List<ObjectFieldContext> ELMS = new ArrayList<ObjectFieldContext>();
		public List<ObjectFieldContext> objectField() {
			return getRuleContexts(ObjectFieldContext.class);
		}
		public ObjectFieldContext objectField(int i) {
			return getRuleContext(ObjectFieldContext.class,i);
		}
		public ObjectExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class CallExprContext extends ExprContext {
		public ExprContext NAME;
		public FnArgContext fnArg;
		public List<FnArgContext> ARGS = new ArrayList<FnArgContext>();
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public List<FnArgContext> fnArg() {
			return getRuleContexts(FnArgContext.class);
		}
		public FnArgContext fnArg(int i) {
			return getRuleContext(FnArgContext.class,i);
		}
		public CallExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class IntExprContext extends ExprContext {
		public IntTypeContext intType() {
			return getRuleContext(IntTypeContext.class,0);
		}
		public IntExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class NotExprContext extends ExprContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public NotExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class MathExprContext extends ExprContext {
		public ExprContext L;
		public ExprContext R;
		public MathSignContext mathSign() {
			return getRuleContext(MathSignContext.class,0);
		}
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public MathExprContext(ExprContext ctx) { copyFrom(ctx); }
	}

	public final ExprContext expr() throws RecognitionException {
		return expr(0);
	}

	private ExprContext expr(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExprContext _localctx = new ExprContext(_ctx, _parentState);
		ExprContext _prevctx = _localctx;
		int _startState = 24;
		enterRecursionRule(_localctx, 24, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(243);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,26,_ctx) ) {
			case 1:
				{
				_localctx = new GroupExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(199);
				match(T__13);
				setState(200);
				((GroupExprContext)_localctx).EXPR = expr(0);
				setState(201);
				match(T__16);
				}
				break;
			case 2:
				{
				_localctx = new BoolExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(203);
				_la = _input.LA(1);
				if ( !(_la==T__27 || _la==T__28) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
				break;
			case 3:
				{
				_localctx = new NullExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(204);
				match(T__29);
				}
				break;
			case 4:
				{
				_localctx = new IdentExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(205);
				ident();
				}
				break;
			case 5:
				{
				_localctx = new IntExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(206);
				intType();
				}
				break;
			case 6:
				{
				_localctx = new FloatExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(207);
				floatType();
				}
				break;
			case 7:
				{
				_localctx = new StringExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(208);
				stringType();
				}
				break;
			case 8:
				{
				_localctx = new NotExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(209);
				match(T__30);
				setState(210);
				expr(13);
				}
				break;
			case 9:
				{
				_localctx = new NegativeExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(211);
				match(T__1);
				setState(212);
				expr(12);
				}
				break;
			case 10:
				{
				_localctx = new DecExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(213);
				ident();
				setState(214);
				match(T__38);
				}
				break;
			case 11:
				{
				_localctx = new ArrayExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(216);
				match(T__5);
				setState(225);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__1) | (1L << T__5) | (1L << T__13) | (1L << T__17) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << Ident) | (1L << Int) | (1L << Str) | (1L << Float))) != 0)) {
					{
					setState(217);
					((ArrayExprContext)_localctx).expr = expr(0);
					((ArrayExprContext)_localctx).ELMS.add(((ArrayExprContext)_localctx).expr);
					setState(222);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==T__14) {
						{
						{
						setState(218);
						match(T__14);
						setState(219);
						((ArrayExprContext)_localctx).expr = expr(0);
						((ArrayExprContext)_localctx).ELMS.add(((ArrayExprContext)_localctx).expr);
						}
						}
						setState(224);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					}
				}

				setState(227);
				match(T__6);
				}
				break;
			case 12:
				{
				_localctx = new ObjectExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(228);
				match(T__17);
				setState(240);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__5 || _la==Ident) {
					{
					setState(229);
					((ObjectExprContext)_localctx).objectField = objectField();
					((ObjectExprContext)_localctx).ELMS.add(((ObjectExprContext)_localctx).objectField);
					setState(234);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,23,_ctx);
					while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
						if ( _alt==1 ) {
							{
							{
							setState(230);
							match(T__14);
							setState(231);
							((ObjectExprContext)_localctx).objectField = objectField();
							((ObjectExprContext)_localctx).ELMS.add(((ObjectExprContext)_localctx).objectField);
							}
							} 
						}
						setState(236);
						_errHandler.sync(this);
						_alt = getInterpreter().adaptivePredict(_input,23,_ctx);
					}
					setState(238);
					_errHandler.sync(this);
					_la = _input.LA(1);
					if (_la==T__14) {
						{
						setState(237);
						match(T__14);
						}
					}

					}
				}

				setState(242);
				match(T__18);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(292);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,30,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(290);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,29,_ctx) ) {
					case 1:
						{
						_localctx = new MathExprContext(new ExprContext(_parentctx, _parentState));
						((MathExprContext)_localctx).L = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(245);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(246);
						mathSign();
						setState(247);
						((MathExprContext)_localctx).R = expr(12);
						}
						break;
					case 2:
						{
						_localctx = new EqualExprContext(new ExprContext(_parentctx, _parentState));
						((EqualExprContext)_localctx).L = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(249);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(250);
						match(T__31);
						setState(251);
						((EqualExprContext)_localctx).R = expr(11);
						}
						break;
					case 3:
						{
						_localctx = new NotEqualExprContext(new ExprContext(_parentctx, _parentState));
						((NotEqualExprContext)_localctx).L = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(252);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(253);
						match(T__32);
						setState(254);
						((NotEqualExprContext)_localctx).R = expr(10);
						}
						break;
					case 4:
						{
						_localctx = new GtExprContext(new ExprContext(_parentctx, _parentState));
						((GtExprContext)_localctx).L = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(255);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(256);
						match(T__33);
						setState(257);
						((GtExprContext)_localctx).R = expr(9);
						}
						break;
					case 5:
						{
						_localctx = new LteExprContext(new ExprContext(_parentctx, _parentState));
						((LteExprContext)_localctx).L = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(258);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(259);
						match(T__34);
						setState(260);
						((LteExprContext)_localctx).R = expr(8);
						}
						break;
					case 6:
						{
						_localctx = new LtExprContext(new ExprContext(_parentctx, _parentState));
						((LtExprContext)_localctx).L = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(261);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(262);
						match(T__35);
						setState(263);
						((LtExprContext)_localctx).R = expr(7);
						}
						break;
					case 7:
						{
						_localctx = new GteExprContext(new ExprContext(_parentctx, _parentState));
						((GteExprContext)_localctx).L = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(264);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(265);
						match(T__36);
						setState(266);
						((GteExprContext)_localctx).R = expr(6);
						}
						break;
					case 8:
						{
						_localctx = new CallExprContext(new ExprContext(_parentctx, _parentState));
						((CallExprContext)_localctx).NAME = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(267);
						if (!(precpred(_ctx, 20))) throw new FailedPredicateException(this, "precpred(_ctx, 20)");
						setState(268);
						match(T__13);
						setState(277);
						_errHandler.sync(this);
						_la = _input.LA(1);
						if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__1) | (1L << T__5) | (1L << T__13) | (1L << T__17) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << Ident) | (1L << Int) | (1L << Str) | (1L << Float))) != 0)) {
							{
							setState(269);
							((CallExprContext)_localctx).fnArg = fnArg();
							((CallExprContext)_localctx).ARGS.add(((CallExprContext)_localctx).fnArg);
							setState(274);
							_errHandler.sync(this);
							_la = _input.LA(1);
							while (_la==T__14) {
								{
								{
								setState(270);
								match(T__14);
								setState(271);
								((CallExprContext)_localctx).fnArg = fnArg();
								((CallExprContext)_localctx).ARGS.add(((CallExprContext)_localctx).fnArg);
								}
								}
								setState(276);
								_errHandler.sync(this);
								_la = _input.LA(1);
							}
							}
						}

						setState(279);
						match(T__16);
						}
						break;
					case 9:
						{
						_localctx = new IndexExprContext(new ExprContext(_parentctx, _parentState));
						((IndexExprContext)_localctx).PARENT = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(280);
						if (!(precpred(_ctx, 18))) throw new FailedPredicateException(this, "precpred(_ctx, 18)");
						setState(281);
						match(T__5);
						setState(282);
						((IndexExprContext)_localctx).CHILD = expr(0);
						setState(283);
						match(T__6);
						}
						break;
					case 10:
						{
						_localctx = new DotExprContext(new ExprContext(_parentctx, _parentState));
						((DotExprContext)_localctx).PARENT = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(285);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(286);
						match(T__4);
						setState(287);
						((DotExprContext)_localctx).CHILD = ident();
						}
						break;
					case 11:
						{
						_localctx = new IncExprContext(new ExprContext(_parentctx, _parentState));
						((IncExprContext)_localctx).L = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(288);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(289);
						match(T__37);
						}
						break;
					}
					} 
				}
				setState(294);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,30,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 12:
			return expr_sempred((ExprContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 11);
		case 1:
			return precpred(_ctx, 10);
		case 2:
			return precpred(_ctx, 9);
		case 3:
			return precpred(_ctx, 8);
		case 4:
			return precpred(_ctx, 7);
		case 5:
			return precpred(_ctx, 6);
		case 6:
			return precpred(_ctx, 5);
		case 7:
			return precpred(_ctx, 20);
		case 8:
			return precpred(_ctx, 18);
		case 9:
			return precpred(_ctx, 17);
		case 10:
			return precpred(_ctx, 4);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3.\u012a\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\3\2\3\2\3\2\3\3\7\3!\n\3\f\3\16\3$\13\3"+
		"\3\4\3\4\3\5\3\5\3\6\3\6\3\7\3\7\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\t\5\t\66"+
		"\n\t\3\n\3\n\3\n\3\n\5\n<\n\n\3\n\5\n?\n\n\3\n\3\n\3\n\3\n\3\n\3\n\7\n"+
		"G\n\n\f\n\16\nJ\13\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\7\nV\n\n"+
		"\f\n\16\nY\13\n\3\n\3\n\3\n\5\n^\n\n\3\n\3\n\5\nb\n\n\5\nd\n\n\3\n\3\n"+
		"\3\n\3\n\3\n\3\n\3\n\3\n\3\n\7\no\n\n\f\n\16\nr\13\n\3\n\3\n\3\n\3\n\3"+
		"\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\5\n\u0084\n\n\3\n\3\n\3"+
		"\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\7\n\u0092\n\n\f\n\16\n\u0095\13"+
		"\n\3\n\3\n\3\n\3\n\3\n\5\n\u009c\n\n\3\n\3\n\5\n\u00a0\n\n\3\n\3\n\5\n"+
		"\u00a4\n\n\3\13\3\13\3\13\3\13\5\13\u00aa\n\13\3\f\3\f\5\f\u00ae\n\f\3"+
		"\f\3\f\3\f\3\f\7\f\u00b4\n\f\f\f\16\f\u00b7\13\f\5\f\u00b9\n\f\3\f\3\f"+
		"\5\f\u00bd\n\f\3\r\3\r\3\r\3\r\3\r\5\r\u00c4\n\r\3\r\3\r\3\r\3\16\3\16"+
		"\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16"+
		"\3\16\3\16\3\16\3\16\3\16\3\16\7\16\u00df\n\16\f\16\16\16\u00e2\13\16"+
		"\5\16\u00e4\n\16\3\16\3\16\3\16\3\16\3\16\7\16\u00eb\n\16\f\16\16\16\u00ee"+
		"\13\16\3\16\5\16\u00f1\n\16\5\16\u00f3\n\16\3\16\5\16\u00f6\n\16\3\16"+
		"\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16"+
		"\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\7\16\u0113"+
		"\n\16\f\16\16\16\u0116\13\16\5\16\u0118\n\16\3\16\3\16\3\16\3\16\3\16"+
		"\3\16\3\16\3\16\3\16\3\16\3\16\7\16\u0125\n\16\f\16\16\16\u0128\13\16"+
		"\3\16\2\3\32\17\2\4\6\b\n\f\16\20\22\24\26\30\32\2\4\3\2\3\6\3\2\36\37"+
		"\2\u0157\2\34\3\2\2\2\4\"\3\2\2\2\6%\3\2\2\2\b\'\3\2\2\2\n)\3\2\2\2\f"+
		"+\3\2\2\2\16-\3\2\2\2\20\65\3\2\2\2\22\u00a3\3\2\2\2\24\u00a9\3\2\2\2"+
		"\26\u00ad\3\2\2\2\30\u00c3\3\2\2\2\32\u00f5\3\2\2\2\34\35\5\4\3\2\35\36"+
		"\7\2\2\3\36\3\3\2\2\2\37!\5\22\n\2 \37\3\2\2\2!$\3\2\2\2\" \3\2\2\2\""+
		"#\3\2\2\2#\5\3\2\2\2$\"\3\2\2\2%&\7*\2\2&\7\3\2\2\2\'(\7+\2\2(\t\3\2\2"+
		"\2)*\7-\2\2*\13\3\2\2\2+,\7,\2\2,\r\3\2\2\2-.\t\2\2\2.\17\3\2\2\2/\60"+
		"\7\7\2\2\60\66\5\6\4\2\61\62\7\b\2\2\62\63\5\32\16\2\63\64\7\t\2\2\64"+
		"\66\3\2\2\2\65/\3\2\2\2\65\61\3\2\2\2\66\21\3\2\2\2\678\7\n\2\28;\5\f"+
		"\7\29:\7\13\2\2:<\5\6\4\2;9\3\2\2\2;<\3\2\2\2<\u00a4\3\2\2\2=?\7\f\2\2"+
		">=\3\2\2\2>?\3\2\2\2?@\3\2\2\2@A\5\6\4\2AB\7\r\2\2BC\5\32\16\2C\u00a4"+
		"\3\2\2\2DH\5\6\4\2EG\5\20\t\2FE\3\2\2\2GJ\3\2\2\2HF\3\2\2\2HI\3\2\2\2"+
		"IK\3\2\2\2JH\3\2\2\2KL\7\16\2\2LM\5\32\16\2M\u00a4\3\2\2\2N\u00a4\5\26"+
		"\f\2OP\7\17\2\2PQ\5\6\4\2Qc\7\20\2\2RW\5\6\4\2ST\7\21\2\2TV\5\6\4\2US"+
		"\3\2\2\2VY\3\2\2\2WU\3\2\2\2WX\3\2\2\2X]\3\2\2\2YW\3\2\2\2Z[\7\21\2\2"+
		"[\\\7\22\2\2\\^\5\6\4\2]Z\3\2\2\2]^\3\2\2\2^b\3\2\2\2_`\7\22\2\2`b\5\6"+
		"\4\2aR\3\2\2\2a_\3\2\2\2bd\3\2\2\2ca\3\2\2\2cd\3\2\2\2de\3\2\2\2ef\7\23"+
		"\2\2fg\7\24\2\2gh\5\4\3\2hi\7\25\2\2i\u00a4\3\2\2\2jk\7\26\2\2kp\5\6\4"+
		"\2lm\7\21\2\2mo\5\6\4\2nl\3\2\2\2or\3\2\2\2pn\3\2\2\2pq\3\2\2\2qs\3\2"+
		"\2\2rp\3\2\2\2st\7\27\2\2tu\5\32\16\2uv\7\24\2\2vw\5\4\3\2wx\7\25\2\2"+
		"x\u00a4\3\2\2\2yz\7\26\2\2z{\5\32\16\2{|\7\24\2\2|}\5\4\3\2}\u0083\7\25"+
		"\2\2~\177\7\30\2\2\177\u0080\7\24\2\2\u0080\u0081\5\4\3\2\u0081\u0082"+
		"\7\25\2\2\u0082\u0084\3\2\2\2\u0083~\3\2\2\2\u0083\u0084\3\2\2\2\u0084"+
		"\u00a4\3\2\2\2\u0085\u0086\7\31\2\2\u0086\u0087\5\32\16\2\u0087\u0088"+
		"\7\24\2\2\u0088\u0089\5\4\3\2\u0089\u0093\7\25\2\2\u008a\u008b\7\30\2"+
		"\2\u008b\u008c\7\31\2\2\u008c\u008d\5\32\16\2\u008d\u008e\7\24\2\2\u008e"+
		"\u008f\5\4\3\2\u008f\u0090\7\25\2\2\u0090\u0092\3\2\2\2\u0091\u008a\3"+
		"\2\2\2\u0092\u0095\3\2\2\2\u0093\u0091\3\2\2\2\u0093\u0094\3\2\2\2\u0094"+
		"\u009b\3\2\2\2\u0095\u0093\3\2\2\2\u0096\u0097\7\30\2\2\u0097\u0098\7"+
		"\24\2\2\u0098\u0099\5\4\3\2\u0099\u009a\7\25\2\2\u009a\u009c\3\2\2\2\u009b"+
		"\u0096\3\2\2\2\u009b\u009c\3\2\2\2\u009c\u00a4\3\2\2\2\u009d\u009f\7\32"+
		"\2\2\u009e\u00a0\5\32\16\2\u009f\u009e\3\2\2\2\u009f\u00a0\3\2\2\2\u00a0"+
		"\u00a4\3\2\2\2\u00a1\u00a4\7\33\2\2\u00a2\u00a4\5\32\16\2\u00a3\67\3\2"+
		"\2\2\u00a3>\3\2\2\2\u00a3D\3\2\2\2\u00a3N\3\2\2\2\u00a3O\3\2\2\2\u00a3"+
		"j\3\2\2\2\u00a3y\3\2\2\2\u00a3\u0085\3\2\2\2\u00a3\u009d\3\2\2\2\u00a3"+
		"\u00a1\3\2\2\2\u00a3\u00a2\3\2\2\2\u00a4\23\3\2\2\2\u00a5\u00aa\5\32\16"+
		"\2\u00a6\u00a7\5\32\16\2\u00a7\u00a8\7\22\2\2\u00a8\u00aa\3\2\2\2\u00a9"+
		"\u00a5\3\2\2\2\u00a9\u00a6\3\2\2\2\u00aa\25\3\2\2\2\u00ab\u00ae\5\6\4"+
		"\2\u00ac\u00ae\5\32\16\2\u00ad\u00ab\3\2\2\2\u00ad\u00ac\3\2\2\2\u00ae"+
		"\u00af\3\2\2\2\u00af\u00b8\7\20\2\2\u00b0\u00b5\5\24\13\2\u00b1\u00b2"+
		"\7\21\2\2\u00b2\u00b4\5\24\13\2\u00b3\u00b1\3\2\2\2\u00b4\u00b7\3\2\2"+
		"\2\u00b5\u00b3\3\2\2\2\u00b5\u00b6\3\2\2\2\u00b6\u00b9\3\2\2\2\u00b7\u00b5"+
		"\3\2\2\2\u00b8\u00b0\3\2\2\2\u00b8\u00b9\3\2\2\2\u00b9\u00ba\3\2\2\2\u00ba"+
		"\u00bc\7\23\2\2\u00bb\u00bd\7\34\2\2\u00bc\u00bb\3\2\2\2\u00bc\u00bd\3"+
		"\2\2\2\u00bd\27\3\2\2\2\u00be\u00c4\5\6\4\2\u00bf\u00c0\7\b\2\2\u00c0"+
		"\u00c1\5\32\16\2\u00c1\u00c2\7\t\2\2\u00c2\u00c4\3\2\2\2\u00c3\u00be\3"+
		"\2\2\2\u00c3\u00bf\3\2\2\2\u00c4\u00c5\3\2\2\2\u00c5\u00c6\7\35\2\2\u00c6"+
		"\u00c7\5\32\16\2\u00c7\31\3\2\2\2\u00c8\u00c9\b\16\1\2\u00c9\u00ca\7\20"+
		"\2\2\u00ca\u00cb\5\32\16\2\u00cb\u00cc\7\23\2\2\u00cc\u00f6\3\2\2\2\u00cd"+
		"\u00f6\t\3\2\2\u00ce\u00f6\7 \2\2\u00cf\u00f6\5\6\4\2\u00d0\u00f6\5\b"+
		"\5\2\u00d1\u00f6\5\n\6\2\u00d2\u00f6\5\f\7\2\u00d3\u00d4\7!\2\2\u00d4"+
		"\u00f6\5\32\16\17\u00d5\u00d6\7\4\2\2\u00d6\u00f6\5\32\16\16\u00d7\u00d8"+
		"\5\6\4\2\u00d8\u00d9\7)\2\2\u00d9\u00f6\3\2\2\2\u00da\u00e3\7\b\2\2\u00db"+
		"\u00e0\5\32\16\2\u00dc\u00dd\7\21\2\2\u00dd\u00df\5\32\16\2\u00de\u00dc"+
		"\3\2\2\2\u00df\u00e2\3\2\2\2\u00e0\u00de\3\2\2\2\u00e0\u00e1\3\2\2\2\u00e1"+
		"\u00e4\3\2\2\2\u00e2\u00e0\3\2\2\2\u00e3\u00db\3\2\2\2\u00e3\u00e4\3\2"+
		"\2\2\u00e4\u00e5\3\2\2\2\u00e5\u00f6\7\t\2\2\u00e6\u00f2\7\24\2\2\u00e7"+
		"\u00ec\5\30\r\2\u00e8\u00e9\7\21\2\2\u00e9\u00eb\5\30\r\2\u00ea\u00e8"+
		"\3\2\2\2\u00eb\u00ee\3\2\2\2\u00ec\u00ea\3\2\2\2\u00ec\u00ed\3\2\2\2\u00ed"+
		"\u00f0\3\2\2\2\u00ee\u00ec\3\2\2\2\u00ef\u00f1\7\21\2\2\u00f0\u00ef\3"+
		"\2\2\2\u00f0\u00f1\3\2\2\2\u00f1\u00f3\3\2\2\2\u00f2\u00e7\3\2\2\2\u00f2"+
		"\u00f3\3\2\2\2\u00f3\u00f4\3\2\2\2\u00f4\u00f6\7\25\2\2\u00f5\u00c8\3"+
		"\2\2\2\u00f5\u00cd\3\2\2\2\u00f5\u00ce\3\2\2\2\u00f5\u00cf\3\2\2\2\u00f5"+
		"\u00d0\3\2\2\2\u00f5\u00d1\3\2\2\2\u00f5\u00d2\3\2\2\2\u00f5\u00d3\3\2"+
		"\2\2\u00f5\u00d5\3\2\2\2\u00f5\u00d7\3\2\2\2\u00f5\u00da\3\2\2\2\u00f5"+
		"\u00e6\3\2\2\2\u00f6\u0126\3\2\2\2\u00f7\u00f8\f\r\2\2\u00f8\u00f9\5\16"+
		"\b\2\u00f9\u00fa\5\32\16\16\u00fa\u0125\3\2\2\2\u00fb\u00fc\f\f\2\2\u00fc"+
		"\u00fd\7\"\2\2\u00fd\u0125\5\32\16\r\u00fe\u00ff\f\13\2\2\u00ff\u0100"+
		"\7#\2\2\u0100\u0125\5\32\16\f\u0101\u0102\f\n\2\2\u0102\u0103\7$\2\2\u0103"+
		"\u0125\5\32\16\13\u0104\u0105\f\t\2\2\u0105\u0106\7%\2\2\u0106\u0125\5"+
		"\32\16\n\u0107\u0108\f\b\2\2\u0108\u0109\7&\2\2\u0109\u0125\5\32\16\t"+
		"\u010a\u010b\f\7\2\2\u010b\u010c\7\'\2\2\u010c\u0125\5\32\16\b\u010d\u010e"+
		"\f\26\2\2\u010e\u0117\7\20\2\2\u010f\u0114\5\24\13\2\u0110\u0111\7\21"+
		"\2\2\u0111\u0113\5\24\13\2\u0112\u0110\3\2\2\2\u0113\u0116\3\2\2\2\u0114"+
		"\u0112\3\2\2\2\u0114\u0115\3\2\2\2\u0115\u0118\3\2\2\2\u0116\u0114\3\2"+
		"\2\2\u0117\u010f\3\2\2\2\u0117\u0118\3\2\2\2\u0118\u0119\3\2\2\2\u0119"+
		"\u0125\7\23\2\2\u011a\u011b\f\24\2\2\u011b\u011c\7\b\2\2\u011c\u011d\5"+
		"\32\16\2\u011d\u011e\7\t\2\2\u011e\u0125\3\2\2\2\u011f\u0120\f\23\2\2"+
		"\u0120\u0121\7\7\2\2\u0121\u0125\5\6\4\2\u0122\u0123\f\6\2\2\u0123\u0125"+
		"\7(\2\2\u0124\u00f7\3\2\2\2\u0124\u00fb\3\2\2\2\u0124\u00fe\3\2\2\2\u0124"+
		"\u0101\3\2\2\2\u0124\u0104\3\2\2\2\u0124\u0107\3\2\2\2\u0124\u010a\3\2"+
		"\2\2\u0124\u010d\3\2\2\2\u0124\u011a\3\2\2\2\u0124\u011f\3\2\2\2\u0124"+
		"\u0122\3\2\2\2\u0125\u0128\3\2\2\2\u0126\u0124\3\2\2\2\u0126\u0127\3\2"+
		"\2\2\u0127\33\3\2\2\2\u0128\u0126\3\2\2\2!\"\65;>HW]acp\u0083\u0093\u009b"+
		"\u009f\u00a3\u00a9\u00ad\u00b5\u00b8\u00bc\u00c3\u00e0\u00e3\u00ec\u00f0"+
		"\u00f2\u00f5\u0114\u0117\u0124\u0126";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}